package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFailed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderFailed0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path stroke-linecap="round" d="m19 22l10 10m0-10L19 32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderFailed0)"/>`),
		g.Group(children),
	)
}