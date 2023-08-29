package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zm0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zm0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zm4-2h7c.6 0 1-.4 1-1s-.4-1-1-1h-7c-.6 0-1 .4-1 1s.4 1 1 1zm-4 14H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zm11-4h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zm0-8h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zM10 3H7c-.6 0-1 .4-1 1s.4 1 1 1h3c.6 0 1-.4 1-1s-.4-1-1-1zm11 8h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1zm-4 8h-3c-.6 0-1 .4-1 1s.4 1 1 1h3c.6 0 1-.4 1-1s-.4-1-1-1z"/>`),
		g.Group(children),
	)
}