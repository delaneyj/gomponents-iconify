package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneSwitchLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 8.62v6.76L5.12 12L8.5 8.62" opacity=".3"/><path fill="currentColor" d="M8.5 8.62v6.76L5.12 12L8.5 8.62M10 5l-7 7l7 7V5zm4 0v14l7-7l-7-7z"/>`),
		g.Group(children),
	)
}