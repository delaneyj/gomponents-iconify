package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationOctagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.7 7.6l-5.3-5.3c-.2-.2-.4-.3-.7-.3H8.3c-.3 0-.5.1-.7.3L2.3 7.6c-.2.2-.3.4-.3.7v7.5c0 .3.1.5.3.7l5.3 5.3c.2.1.4.2.7.2h7.5c.3 0 .5-.1.7-.3l5.3-5.3c.2-.2.3-.4.3-.7V8.3c-.1-.3-.2-.5-.4-.7zM12 17c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm1-5c0 .6-.4 1-1 1s-1-.4-1-1V8c0-.6.4-1 1-1s1 .4 1 1v4z"/>`),
		g.Group(children),
	)
}