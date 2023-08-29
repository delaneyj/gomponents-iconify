package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#37474F" d="M8 41V7c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H12c-2.2 0-4-1.8-4-4z"/><path fill="#BBDEFB" d="M36 6H12c-.6 0-1 .4-1 1v31c0 .6.4 1 1 1h24c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1z"/><path fill="#78909C" d="M21 41h6v2h-6z"/>`),
		g.Group(children),
	)
}