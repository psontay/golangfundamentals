package practice

import "time"

// the mission of this pattern is construct struct complex, specially when some struct have so much parameter

type Member struct {
	Name     string
	HasPT    bool
	Duration int
}

// functional options
type MemberOption func(*Member)

func WithPT() MemberOption {
	return func(m *Member) {
		m.HasPT = true
	}
}

func WithDuration(months int) MemberOption {
	return func(m *Member) {
		m.Duration = months
	}
}

// constructor
func NewMember(name string, options ...MemberOption) *Member {
	member := &Member{
		Name:     name,
		HasPT:    false,
		Duration: 1,
	}
	for _, option := range options {
		option(member)
	}
	return member
}

// practice

type DBConfig struct {
	Host         string
	Port         int
	Timeout      time.Duration
	MaxIdleConns int
}

type DBOption func(*DBConfig)

func WithTimeOut(timeout time.Duration) DBOption {
	return func(db *DBConfig) {
		db.Timeout = timeout
	}
}

func WithMaxIdle(conns int) DBOption {
	return func(db *DBConfig) {
		db.MaxIdleConns = conns
	}
}

func NewDBConfig(host string, port int, opts ...DBOption) *DBConfig {
	dbConfig := &DBConfig{
		Host:         host,
		Port:         port,
		Timeout:      30 * time.Second,
		MaxIdleConns: 5,
	}
	for _, option := range opts {
		option(dbConfig)
	}
	return dbConfig
}
