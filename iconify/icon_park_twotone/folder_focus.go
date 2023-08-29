package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderFocus0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path stroke-linecap="round" d="m24 20l2.243 4.913l5.366.615l-3.98 3.651l1.073 5.293L24 31.816l-4.702 2.656l1.073-5.293l-3.98-3.651l5.366-.615L24 20Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderFocus0)"/>`),
		g.Group(children),
	)
}