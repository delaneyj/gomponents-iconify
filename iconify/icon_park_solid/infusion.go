package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infusion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInfusion0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M38 30c0 7.732-6.268 14-14 14s-14-6.268-14-14S24 4 24 4s14 18.268 14 26Z"/><path stroke="#000" stroke-linecap="round" d="M18 30h12m-6-6v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInfusion0)"/>`),
		g.Group(children),
	)
}