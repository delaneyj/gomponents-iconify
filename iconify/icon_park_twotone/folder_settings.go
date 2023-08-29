package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderSettings0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><circle cx="24" cy="28" r="4" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 21v3m0 8v3m4.828-12l-2.121 2.121M20.828 31l-2.121 2.121M19 23l2.121 2.121M27 31l2.121 2.121M17 28h3m8 0h3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderSettings0)"/>`),
		g.Group(children),
	)
}