package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RetroBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRetroBag0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M35 14c0-5.523-4.925-10-11-10S13 8.477 13 14"/><path fill="#555" d="M7 16a2 2 0 0 1 2-2h30a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-9a2 2 0 0 1-2-2v0a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v0a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V16Z"/><path d="M10 30v14h28V30"/><path d="M20 26h8v6h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRetroBag0)"/>`),
		g.Group(children),
	)
}