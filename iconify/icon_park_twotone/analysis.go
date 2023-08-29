package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAnalysis0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M44 5H4v12h40V5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m4 41.03l12.176-12.3l6.579 6.3L30.798 27l4.48 4.368"/><path stroke-linecap="round" d="M44 16.172v26m-40-26v14M13.015 43H44M17 11h21m-28-.003h1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAnalysis0)"/>`),
		g.Group(children),
	)
}