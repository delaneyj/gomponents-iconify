package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSystemOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 8h-2.59l-1-1H15a2.006 2.006 0 0 0-2 2v4a2.006 2.006 0 0 0 2 2h6a2.006 2.006 0 0 0 2-2v-3a2.006 2.006 0 0 0-2-2Zm0 5h-6V9h2l1 1h3Zm0 4h-2.59l-1-1H15a2.006 2.006 0 0 0-2 2v4a2.006 2.006 0 0 0 2 2h6a2.006 2.006 0 0 0 2-2v-3a2.006 2.006 0 0 0-2-2Zm0 5h-6v-4h2l1 1h3ZM8 19v-7h5v-2H8V8H6v13h7v-2H8zm2-18H7.41l-1-1H4a2.006 2.006 0 0 0-2 2v4a2.006 2.006 0 0 0 2 2h6a2.006 2.006 0 0 0 2-2V3a2.006 2.006 0 0 0-2-2Zm0 5H4V2h2l1 1h3Z"/>`),
		g.Group(children),
	)
}