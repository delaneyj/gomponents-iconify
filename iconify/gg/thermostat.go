package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermostat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M15 14a5 5 0 1 1-6 0V4a3 3 0 1 1 6 0v10ZM13 4v11.17A3.001 3.001 0 0 1 12 21a3 3 0 0 1-1-5.83V4a1 1 0 1 1 2 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}