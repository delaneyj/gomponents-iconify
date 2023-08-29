package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YandexInternational(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20.8 1l-5.6 16.2l-5-13.2H7l7 18.6V31h3v-9.9L24 1h-3.2z"/>`),
		g.Group(children),
	)
}