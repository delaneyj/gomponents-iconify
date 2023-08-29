package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.11 21.46l-1.27 1.27L18.11 20H4a2 2 0 0 1-2-2V6c0-.58.25-1.1.64-1.47L1.11 3l1.28-1.27l19.72 19.73M22 18V8a2 2 0 0 0-2-2h-8l-2-2H7.2l14.68 14.68c.08-.21.12-.44.12-.68Z"/>`),
		g.Group(children),
	)
}