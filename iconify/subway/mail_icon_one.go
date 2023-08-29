package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailIconOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M291.1 147.9V41.3L120.5 212l170.7 170.7V284c96.4 6.7 165 39.4 220.8 199c0-90.8 8.2-296.7-220.9-335.1zm-120.4-26.3V41.3L0 212l170.7 170.7v-80.3L80.3 212l90.4-90.4z"/>`),
		g.Group(children),
	)
}