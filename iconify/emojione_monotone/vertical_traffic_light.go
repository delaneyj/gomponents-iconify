package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2c-8.25 0-15 6.751-15 15v30c0 8.25 6.75 15 15 15s15-6.75 15-15V17c0-8.249-6.75-15-15-15m0 56.25a7.5 7.5 0 0 1-7.5-7.5a7.5 7.5 0 1 1 7.5 7.5m0-18.75a7.5 7.5 0 0 1-7.5-7.5a7.5 7.5 0 1 1 7.5 7.5m0-18.75a7.5 7.5 0 0 1-7.5-7.5a7.5 7.5 0 1 1 7.5 7.5"/>`),
		g.Group(children),
	)
}