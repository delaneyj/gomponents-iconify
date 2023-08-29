package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailIconTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M341.3 32.7V113l90.3 90.3l-90.3 90.3V374L512 203.3L341.3 32.7zm50.2 170.6L220.9 32.7v106.6C-8.2 177.6 0 383.5 0 474.4c55.9-159.6 124.5-192.3 220.8-199V374l170.7-170.7z"/>`),
		g.Group(children),
	)
}