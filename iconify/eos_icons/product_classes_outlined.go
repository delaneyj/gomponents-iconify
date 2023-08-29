package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductClassesOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="5" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M4 4h2v9H4z"/><path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm0 19H3V3h4Zm7-18v7h-3V3h3m1-1h-5v9h5V2Zm6 1v7h-3V3h3m1-1h-5v9h5V2Zm-8 12v7h-3v-7h3m1-1h-5v9h5v-9Zm6 1v7h-3v-7h3m1-1h-5v9h5v-9Z"/>`),
		g.Group(children),
	)
}