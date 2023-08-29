package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiddleBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#CFD8DC"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2z"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1z"/></g><path fill="#8BC34A" d="M34 44H14c-1.1 0-2-.9-2-2V23h24v19c0 1.1-.9 2-2 2z"/>`),
		g.Group(children),
	)
}