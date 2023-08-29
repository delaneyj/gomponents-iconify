package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 6.5h-3v-4c0-.6-.4-1-1-1h-11c-.6 0-1 .4-1 1v4h-3c-.6 0-1 .4-1 1v14c0 .6.4 1 1 1h19c.6 0 1-.4 1-1v-14c0-.6-.4-1-1-1zm-14 12h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1zm0-4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1zm5 4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1zm0-4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1zm1-5.5H13v.5c0 .6-.4 1-1 1s-1-.4-1-1V9h-.5c-.6 0-1-.4-1-1s.4-1 1-1h.5v-.5c0-.6.4-1 1-1s1 .4 1 1V7h.5c.6 0 1 .4 1 1s-.4 1-1 1zm4 9.5h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1zm0-4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1z"/>`),
		g.Group(children),
	)
}