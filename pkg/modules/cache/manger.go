package cache

import (
	"github.com/go-redis/redis/v8"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

var manager *Manager

type Manager struct {
	mutex sync.Mutex

	RedisConnections map[string]*redisClientHolder
}

type redisClientHolder struct {
	redis.UniversalClient
	name  []string
	count int64
}

func GetManager() *Manager {
	if manager == nil {
		manager = &Manager{
			RedisConnections: make(map[string]*redisClientHolder),
		}
	}
	return manager
}

// GetRedisClient gets a redis client for a particular connection
func (m *Manager) GetRedisClient(connection string) redis.UniversalClient {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	client, ok := m.RedisConnections[connection]
	if ok {
		client.count++
		return client
	}
	uri := ToRedisURI(connection)
	client = &redisClientHolder{
		name: []string{connection, uri.String()},
	}

	opts := &redis.UniversalOptions{}
	//tlsConfig := &tls.Config{}

	// Handle username/password
	if password, ok := uri.User.Password(); ok {
		opts.Password = password
		// Username does not appear to be handled by redis.Options
		opts.Username = uri.User.Username()
	} else if uri.User.Username() != "" {
		// assume this is the password
		opts.Password = uri.User.Username()
	}

	if uri.Host != "" {
		opts.Addrs = append(opts.Addrs, strings.Split(uri.Host, ",")...)
	}
	if uri.Path != "" {
		if db, err := strconv.Atoi(uri.Path); err == nil {
			opts.DB = db
		}
	}
	client.UniversalClient = redis.NewClient(opts.Simple())
	for _, name := range client.name {
		m.RedisConnections[name] = client
	}

	client.count++

	return client
}

func ToRedisURI(connection string) *url.URL {
	uri, err := url.Parse(connection)
	if err == nil && strings.HasPrefix(uri.Scheme, "redis") {
		// OK we're going to assume that this is a reasonable redis URI
		return uri
	}

	// Let's set a nice default
	uri, _ = url.Parse("redis://127.0.0.1:6379/0")
	network := "tcp"
	query := uri.Query()

	// OK so there are two types: Space delimited and Comma delimited
	// Let's assume that we have a space delimited string - as this is the most common
	fields := strings.Fields(connection)
	if len(fields) == 1 {
		// It's a comma delimited string, then...
		fields = strings.Split(connection, ",")

	}
	for _, f := range fields {
		items := strings.SplitN(f, "=", 2)
		if len(items) < 2 {
			continue
		}
		switch strings.ToLower(items[0]) {
		case "network":
			if items[1] == "unix" {
				uri.Scheme = "redis+socket"
			}
			network = items[1]
		case "addrs":
			uri.Host = items[1]
			// now we need to handle the clustering
			if strings.Contains(items[1], ",") && network == "tcp" {
				uri.Scheme = "redis+cluster"
			}
		case "addr":
			uri.Host = items[1]
		case "password":
			uri.User = url.UserPassword(uri.User.Username(), items[1])
		case "username":
			password, set := uri.User.Password()
			if !set {
				uri.User = url.User(items[1])
			} else {
				uri.User = url.UserPassword(items[1], password)
			}
		case "db":
			uri.Path = "/" + items[1]
		case "idle_timeout":
			_, err := strconv.Atoi(items[1])
			if err == nil {
				query.Add("idle_timeout", items[1]+"s")
			} else {
				query.Add("idle_timeout", items[1])
			}
		default:
			// Other options become query params
			query.Add(items[0], items[1])
		}
	}

	// Finally we need to fix up the Host if we have a unix port
	if uri.Scheme == "redis+socket" {
		query.Set("db", uri.Path)
		uri.Path = uri.Host
		uri.Host = ""
	}
	uri.RawQuery = query.Encode()

	return uri
}
