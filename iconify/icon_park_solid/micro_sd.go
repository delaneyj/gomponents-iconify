package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMicroSd0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M13 7a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v34a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2v-5l5-4v-4H8v-3l5-6V7Z"/><path stroke="#000" d="M32 11v4m-12-4v4m6-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMicroSd0)"/>`),
		g.Group(children),
	)
}