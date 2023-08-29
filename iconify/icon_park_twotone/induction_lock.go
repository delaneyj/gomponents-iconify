package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InductionLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInductionLock0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M9 18v24h30V18L24 6L9 18Z"/><path d="M24 22v12m7-8v4m-14-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInductionLock0)"/>`),
		g.Group(children),
	)
}