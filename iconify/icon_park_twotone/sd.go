package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSd0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M15.402 5.683A2 2 0 0 1 16.907 5H37a2 2 0 0 1 2 2v34a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V24h4v-6H9v-5l6.402-7.317Z"/><path d="M33 11v4m-12-4v4m6-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSd0)"/>`),
		g.Group(children),
	)
}