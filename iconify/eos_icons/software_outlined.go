package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftwareOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18.315A7.037 7.037 0 0 1 11.263 16H3V4h18v2.264a7.046 7.046 0 0 1 2 2.15V4a2.006 2.006 0 0 0-2-2H3a2.006 2.006 0 0 0-2 2v12a2.006 2.006 0 0 0 2 2h7v2H8v2h8v-2h-2Z"/><path fill="currentColor" d="M17 6a6 6 0 1 0 6 6a5.998 5.998 0 0 0-6-6Zm0 10a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><circle cx="17" cy="12" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}