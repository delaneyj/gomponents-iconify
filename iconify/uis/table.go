package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1zM10 20v-4h4v4h-4zm0-6v-4h4v4h-4zm-6-4h4v4H4v-4zm6-2V4h4v4h-4zm6 2h4v4h-4v-4zm4-2h-4V4h4v4zM8 4v4H4V4h4zM4 16h4v4H4v-4zm12 4v-4h4v4h-4z"/>`),
		g.Group(children),
	)
}