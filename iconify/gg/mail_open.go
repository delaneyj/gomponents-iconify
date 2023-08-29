package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.05 10.015a2 2 0 0 1 .465-2.1L9.879 1.55a3 3 0 0 1 4.242 0l6.364 6.364a2 2 0 0 1 .465 2.101c.032.099.05.204.05.313v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-11a1 1 0 0 1 .05-.313Zm1.879-.687l6.364-6.363a1 1 0 0 1 1.414 0l6.364 6.363h-.03v.03l-6.334 6.334a1 1 0 0 1-1.414 0L4.929 9.328Zm14.07 2.9l-4.878 4.879a3 3 0 0 1-4.242 0l-4.88-4.88v9.101h14v-9.1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}