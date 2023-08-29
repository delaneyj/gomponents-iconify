package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 3.5c0-.6-.4-1-1-1h-16c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1s1-.4 1-1v-15h15c.6 0 1-.4 1-1zm-1 7c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm0 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm-12 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1z"/>`),
		g.Group(children),
	)
}