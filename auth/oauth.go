package auth // import "github.com/lakexyde/wc-api-go/auth"

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/lakexyde/wc-api-go/options"
)

const hashAlgorithm string = "SHA256"

// MicroTimerInterface should return string with microtime which will be used for hashing
type MicroTimerInterface interface {
	Get() string
}

// OAuth authentication for doing non-SSL requests
type OAuth struct {
	URL        string
	Key        string
	Secret     string
	Version    string
	Method     string
	Parameters url.Values
	Timestamp  string
	MicroTimer MicroTimerInterface
}

// GetEnrichedQuery which appended OAuth specific ones
func (o *OAuth) GetEnrichedQuery() url.Values {
	o.Parameters.Set("oauth_consumer_key", o.Key)
	o.Parameters.Set("oauth_timestamp", o.Timestamp)
	o.Parameters.Set("oauth_nonce", oAuthNonce(o.MicroTimer.Get()))
	o.Parameters.Set("oauth_signature_method", "HMAC-"+hashAlgorithm)
	generateSign(o)

	return o.Parameters
}

func generateSign(o *OAuth) {

	o.Parameters = normalize(&o.Parameters)

	qs := queryString(&o.Parameters)

	stringToSign := strings.Join([]string{o.Method, encode(o.URL), encode(qs)}, "&")

	mac := hmac.New(sha256.New, []byte(o.getSecret()))
	mac.Write([]byte(stringToSign))

	var sign string
	sign = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	o.Parameters.Set("oauth_signature", sign)
}

func encode(s string) string {
	s = url.QueryEscape(s)
	// s = strings.ReplaceAll(s, "+", "%20")
	// s = strings.ReplaceAll(s, "%2B", "%20")
	return s
}

func normalize(p *url.Values) url.Values {
	result := url.Values{}
	for k := range *p {
		value := encode(p.Get(k))
		result.Del(k)
		result.Set(encode(k), value)
	}
	return result
}

func queryString(p *url.Values) string {
	var keys []string
	for k := range *p {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var result []string
	for _, key := range keys {
		val := encode(p.Get(key))
		result = append(result, fmt.Sprintf("%s=%s", key, val))
	}
	return strings.Join(result, "&")
}

// SetMicrotimer ...
func (o *OAuth) SetMicrotimer(m MicroTimerInterface) {
	o.MicroTimer = m
}

// SetMethod ...
func (o *OAuth) SetMethod(method string) {
	o.Method = method
}

func (o *OAuth) getSecret() string {
	secret := o.Secret
	if "v1" != o.Version && "v2" != o.Secret {
		secret = secret + "&"
	}
	return secret
}

func oAuthNonce(mc string) string {
	byteHash := sha1.Sum([]byte(mc))
	return fmt.Sprintf("%x", byteHash)
}

// SetOptions ...
func (o *OAuth) SetOptions(opt options.Basic) {
	o.Key = opt.Key
	o.Secret = opt.Secret
	o.Version = opt.Options.Version
	o.Timestamp = opt.OAuthTimestamp()
}

// SetURL ...
func (o *OAuth) SetURL(url string) {
	o.URL = url
}

// SetParameters ...
func (o *OAuth) SetParameters(params url.Values) {
	o.Parameters = params
}
