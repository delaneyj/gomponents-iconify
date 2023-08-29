package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12c0-.6-.4-1-1-1h-7V4c0-.6-.4-1-1-1s-1 .4-1 1v7H4c-.6 0-1 .4-1 1s.4 1 1 1h7v7c0 .6.4 1 1 1s1-.4 1-1v-7h7c.5 0 1-.4 1-1zM4 15c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1zm0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1zM4 9c.6 0 1-.4 1-1s-.5-1-1-1s-1 .4-1 1s.4 1 1 1zm0-4c.6 0 1-.4 1-1s-.5-1-1-1s-1 .4-1 1s.4 1 1 1zm4 0c.6 0 1-.4 1-1s-.5-1-1-1s-1 .4-1 1s.4 1 1 1zm8 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1zM8 19c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1zm8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1zm4-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1zm0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1zm0-10c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1zm0-4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1z"/>`),
		g.Group(children),
	)
}