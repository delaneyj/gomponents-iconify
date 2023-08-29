package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.6 10.9c.1.1.2.1.4.1h7c.6 0 1-.4 1-1c0-.2 0-.3-.1-.4l-3.5-7c-.3-.5-.9-.7-1.4-.4c-.1.1-.3.2-.4.4l-3.5 7c-.2.4 0 1 .5 1.3zm3.9-5.7L19.4 9h-3.8l1.9-3.8zM6.5 2C4 2 2 4 2 6.5S4 11 6.5 11S11 9 11 6.5S9 2 6.5 2zm0 7C5.1 9 4 7.9 4 6.5S5.1 4 6.5 4S9 5.1 9 6.5S7.9 9 6.5 9zm4.2 4.3c-.4-.4-1-.4-1.4 0l-2.8 2.8l-2.8-2.8c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.8 2.8l-2.8 2.8c-.4.4-.4 1 0 1.4s1 .4 1.4 0l2.8-2.8l2.8 2.8c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-2.8-2.8l2.8-2.8c.4-.4.4-1 0-1.4zM21 13h-7c-.6 0-1 .4-1 1v7c0 .6.4 1 1 1h7c.6 0 1-.4 1-1v-7c0-.6-.4-1-1-1zm-1 7h-5v-5h5v5z"/>`),
		g.Group(children),
	)
}