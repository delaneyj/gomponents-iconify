package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YandexInternational(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.2 24v-7.786L0 2.25h2.616L6.45 13.017L10.86-.001h2.405L7.607 16.302v7.697z"/>`),
		g.Group(children),
	)
}