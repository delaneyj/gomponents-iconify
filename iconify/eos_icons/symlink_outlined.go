package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymlinkOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 12l4 3.5l-4 3.5v-2a6.956 6.956 0 0 0-6 3c.56-2.67 2.11-5.46 6-6Zm0-8l5 5h-5Z"/><path fill="currentColor" d="M14.68 0H4a2.006 2.006 0 0 0-2 2v20a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2V7.15ZM20 22H4V2h9.83L20 8Z"/>`),
		g.Group(children),
	)
}