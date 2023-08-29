package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyTaste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBabyTaste0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c9.941 0 18-8.059 18-18S33.941 8 24 8S6 16.059 6 26s8.059 18 18 18Z"/><path stroke-linecap="round" d="M24 8c-.25-1-2-4-6-4m6 4c.083-1 .6-3.2 2-4m5 29s-2 4-7 4s-7-4-7-4"/><path stroke-linecap="round" d="M31 33s1.5-4-1-4s-3 7-3 7m6-15h-4m-12-2v4M4 24v4m40-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBabyTaste0)"/>`),
		g.Group(children),
	)
}