package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm0 2c4 0 7.4 3 7.9 7H12V4zm4 14.9L12.6 13H20c-.4 2.5-1.8 4.7-4 5.9z"/>`),
		g.Group(children),
	)
}