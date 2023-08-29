package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ascending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.08 7.286l.354-.353l.354.353L17 11.498l-.707.708l-3.358-3.359V17h-1V8.847l-3.359 3.359l-.707-.708l4.212-4.212Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}