package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrganizationOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11V9h-2l.025 2H3.5v4h2v-2H11v2h2v-2h5.505l-.005 2h2v-4H13zm-1-9a2.5 2.5 0 1 0 2.5 2.5A2.5 2.5 0 0 0 12 2Zm0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1ZM4.5 17A2.5 2.5 0 1 0 7 19.5A2.5 2.5 0 0 0 4.5 17Zm0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1ZM12 17a2.5 2.5 0 1 0 2.5 2.5A2.5 2.5 0 0 0 12 17Zm0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7.5-3.5a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5Zm0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}