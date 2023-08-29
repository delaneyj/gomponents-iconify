package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M47 17H17C8.75 17 2 23.75 2 32c0 8.251 6.75 15 15 15h30c8.25 0 15-6.749 15-15c0-8.25-6.75-15-15-15M13.25 39.5a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15m18.75 0a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15m18.75 0a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15"/>`),
		g.Group(children),
	)
}