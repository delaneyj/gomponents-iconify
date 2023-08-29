package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 17.5H4v-11h7.8l-.8.8c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l2.5-2.5c.4-.4.4-1 0-1.4l-2.5-2.5c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l.8.8H3c-.6 0-1 .4-1 1v13c0 .6.4 1 1 1h2.5c.6 0 1-.4 1-1s-.4-1-1-1zM21 4.5h-2.5c-.6 0-1 .4-1 1s.4 1 1 1H20v11h-8.4l.8-.8c.4-.4.4-1 0-1.4s-1-.4-1.4 0l-2.5 2.5c-.4.4-.4 1 0 1.4l2.5 2.5c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-.8-.8H21c.6 0 1-.4 1-1v-13c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}