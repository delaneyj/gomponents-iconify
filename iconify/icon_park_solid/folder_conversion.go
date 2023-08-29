package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderConversion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFolderConversion0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path stroke="#000" stroke-linecap="round" d="M17 24h14m-14 6h14m0-6l-5-5m-4 16l-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFolderConversion0)"/>`),
		g.Group(children),
	)
}