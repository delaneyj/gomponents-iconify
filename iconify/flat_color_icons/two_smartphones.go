package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoSmartphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#37474F" d="M6 36V8c0-2.2 1.8-4 4-4h14c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4H10c-2.2 0-4-1.8-4-4z"/><path fill="#BBDEFB" d="M24 7H10c-.6 0-1 .4-1 1v25c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1z"/><path fill="#78909C" d="M14 36h6v2h-6z"/><path fill="#E38939" d="M20 40V12c0-2.2 1.8-4 4-4h14c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4H24c-2.2 0-4-1.8-4-4z"/><path fill="#FFF3E0" d="M38 11H24c-.6 0-1 .4-1 1v25c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V12c0-.6-.4-1-1-1z"/><circle cx="31" cy="41" r="1.5" fill="#A6642A"/>`),
		g.Group(children),
	)
}