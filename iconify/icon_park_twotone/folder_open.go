package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderOpen0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 9v32l5-20h30.5v-6a2 2 0 0 0-2-2H24l-5-6H6a2 2 0 0 0-2 2Z"/><path fill="#555" d="m40 41l4-20H8.812L4 41h36Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderOpen0)"/>`),
		g.Group(children),
	)
}