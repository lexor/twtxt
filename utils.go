package twtxt

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/goware/urlx"
	log "github.com/sirupsen/logrus"
)

func NormalizeURL(url string) string {
	if url == "" {
		return ""
	}
	u, err := urlx.Parse(url)
	if err != nil {
		log.WithError(err).Errorf("error parsing url %s", url)
		return ""
	}
	if u.Scheme == "https" {
		u.Scheme = "http"
		u.Host = strings.TrimSuffix(u.Host, ":443")
	}
	u.User = nil
	u.Path = strings.TrimSuffix(u.Path, "/")
	norm, err := urlx.Normalize(u)
	if err != nil {
		log.WithError(err).Errorf("error normalizing url %s", url)
		return ""
	}
	return norm
}

// FormatMentions turns `@<nick URL>` into `<a href="URL">@nick</a>`
func FormatMentions(text string) template.HTML {
	re := regexp.MustCompile(`@<([^ ]+) *([^>]+)>`)
	return template.HTML(re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		nick, url := parts[1], parts[2]
		return fmt.Sprintf(`<a href="%s">@%s</a>`, url, nick)
	}))
}