package config_test

import (
	"bytes"
	"testing"

	"github.com/prebid/prebid-server/config"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("pbs")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/config")

	viper.SetDefault("external_url", "http://localhost:8000")
	viper.SetDefault("port", 8000)
	viper.SetDefault("admin_port", 6060)
	viper.SetDefault("default_timeout_ms", 250)
	viper.SetDefault("datacache.type", "dummy")

	viper.SetDefault("adapters.pubmatic.endpoint", "http://openbid-useast.pubmatic.com/translator?")
	viper.SetDefault("adapters.rubicon.endpoint", "http://staged-by.rubiconproject.com/a/api/exchange.json")
	viper.SetDefault("adapters.rubicon.usersync_url", "https://pixel.rubiconproject.com/exchange/sync.php?p=prebid")
	viper.SetDefault("adapters.pulsepoint.endpoint", "http://bid.contextweb.com/header/s/ortb/prebid-s2s")
}

func TestDefaults(t *testing.T) {

	cfg, err := config.New()
	if err != nil {
		t.Error(err.Error())
	}

	if cfg.Port != 8000 {
		t.Error("Expected Port 8000")
	}

	if cfg.AdminPort != 6060 {
		t.Error("Expected Admin Port 6060")
	}

	if cfg.DefaultTimeout != uint64(250) {
		t.Error("Expected DefaultTimeout of 250ms")
	}

	if cfg.DataCache.Type != "dummy" {
		t.Error("Expected DataCache Type of 'dummy'")
	}

	if cfg.Adapters["pubmatic"].Endpoint != "http://openbid-useast.pubmatic.com/translator?" {
		t.Errorf("Expected Pubmatic Endpoint of http://openbid-useast.pubmatic.com/translator?")
	}

}

var fullConfig = []byte(`cookie_domain: ".adnxs.com"
external_url: http://prebid.adnxs.com/
host: prebid.adnxs.com
port: 1234
admin_port: 5678
default_timeout_ms: 123
prebid_cache_url: http://prebidcache.net/test/a1?qs=something
require_uuid2: true
recaptcha_secret: asdfasdfasdfasdf
metrics:
  host: upstream:8232
  database: metricsdb
  username: admin
  password: admin1324
datacache:
  type: postgres
  filename: /usr/db/db.db
  dbname: pbsdb
  host: postgres.internal
  user: dbadmin
  password: db2342
  cache_size: 10000000
  ttl_seconds: 3600
adapters:
  indexExchange:
    endpoint: http://ixtest.com/api
  rubicon:
    endpoint: http://rubitest.com/api
    usersync_url: http://pixel.rubiconproject.com/sync.php?p=prebid
    xapi:
      username: rubiuser
      password: rubipw23
  facebook:
    endpoint: http://facebook.com/pbs
    usersync_url: http://facebook.com/ortb/prebid-s2s
    platform_id: abcdefgh1234
`)

func cmpStrings(t *testing.T, key string, a string, b string) {
	if a != b {
		t.Errorf("%s: %s != %s", key, a, b)
	}
}

func cmpInts(t *testing.T, key string, a int, b int) {
	if a != b {
		t.Errorf("%s: %d != %d", key, a, b)
	}
}

func TestFullConfig(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(fullConfig))
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err.Error())
	}
	cmpStrings(t, "cookie domain", cfg.CookieDomain, ".adnxs.com")
	cmpStrings(t, "external url", cfg.ExternalURL, "http://prebid.adnxs.com/")
	cmpStrings(t, "host", cfg.Host, "prebid.adnxs.com")
	cmpInts(t, "port", cfg.Port, 1234)
	cmpInts(t, "admin_port", cfg.AdminPort, 5678)
	if cfg.DefaultTimeout != 123 {
		t.Errorf("DefaultTimeout was %d not 123", cfg.DefaultTimeout)
	}
	cmpStrings(t, "prebid_cache_url", cfg.CacheURL, "http://prebidcache.net/test/a1?qs=something")
	if cfg.RequireUUID2 != true {
		t.Errorf("RequireUUID2 was false")
	}
	cmpStrings(t, "recaptcha_secret", cfg.RecaptchaSecret, "asdfasdfasdfasdf")
	cmpStrings(t, "metrics.host", cfg.Metrics.Host, "upstream:8232")
	cmpStrings(t, "metrics.database", cfg.Metrics.Database, "metricsdb")
	cmpStrings(t, "metrics.username", cfg.Metrics.Username, "admin")
	cmpStrings(t, "metrics.password", cfg.Metrics.Password, "admin1324")
	cmpStrings(t, "datacache.type", cfg.DataCache.Type, "postgres")
	cmpStrings(t, "datacache.filename", cfg.DataCache.Filename, "/usr/db/db.db")
	cmpStrings(t, "datacache.dbname", cfg.DataCache.Database, "pbsdb")
	cmpStrings(t, "datacache.host", cfg.DataCache.Host, "postgres.internal")
	cmpStrings(t, "datacache.user", cfg.DataCache.Username, "dbadmin")
	cmpStrings(t, "datacache.password", cfg.DataCache.Password, "db2342")
	cmpInts(t, "datacache.cache_size", cfg.DataCache.CacheSize, 10000000)
	cmpInts(t, "datacache.ttl_seconds", cfg.DataCache.TTLSeconds, 3600)
	cmpStrings(t, "adapters.indexExchange.endpoint", cfg.Adapters["indexexchange"].Endpoint, "http://ixtest.com/api")
	cmpStrings(t, "adapters.rubicon.endpoint", cfg.Adapters["rubicon"].Endpoint, "http://rubitest.com/api")
	cmpStrings(t, "adapters.rubicon.usersync_url", cfg.Adapters["rubicon"].UserSyncURL, "http://pixel.rubiconproject.com/sync.php?p=prebid")
	cmpStrings(t, "adapters.rubicon.xapi.username", cfg.Adapters["rubicon"].XAPI.Username, "rubiuser")
	cmpStrings(t, "adapters.rubicon.xapi.password", cfg.Adapters["rubicon"].XAPI.Password, "rubipw23")
	cmpStrings(t, "adapters.facebook.endpoint", cfg.Adapters["facebook"].Endpoint, "http://facebook.com/pbs")
	cmpStrings(t, "adapters.facebook.usersync_url", cfg.Adapters["facebook"].UserSyncURL, "http://facebook.com/ortb/prebid-s2s")
	cmpStrings(t, "adapters.facebook.platform_id", cfg.Adapters["facebook"].PlatformID, "abcdefgh1234")
}
