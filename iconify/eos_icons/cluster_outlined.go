package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClusterOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 11H1V1h22ZM3 9h18V3H3Zm20 14H1V13h22ZM3 21h18v-6H3Z"/><path fill="currentColor" d="M4 5h9v2H4z"/><circle cx="16" cy="6" r="1" fill="currentColor"/><circle cx="19" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M4 17h9v2H4z"/><circle cx="16" cy="18" r="1" fill="currentColor"/><circle cx="19" cy="18" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}