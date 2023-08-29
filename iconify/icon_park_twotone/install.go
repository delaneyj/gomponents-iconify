package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Install(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInstall0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M41.4 11.551L36.333 5H11.667l-5.083 6.551"/><path fill="#555" d="M6 13a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v27a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V13Z"/><path stroke-linecap="round" d="m32 27l-8 8l-8-8m7.992-8v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInstall0)"/>`),
		g.Group(children),
	)
}