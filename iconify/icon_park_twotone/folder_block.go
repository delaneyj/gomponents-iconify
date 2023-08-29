package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderBlock0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><circle cx="25" cy="27" r="7" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m27 25l-4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderBlock0)"/>`),
		g.Group(children),
	)
}