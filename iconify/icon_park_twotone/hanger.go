package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHanger0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" d="M20.73 27.125a6 6 0 0 1 6.54 0l15.55 10.108A2.593 2.593 0 0 1 41.407 42H6.593a2.593 2.593 0 0 1-1.413-4.767l15.55-10.107Z" clip-rule="evenodd"/><path d="M24 25s6-8.686 6-12a6 6 0 0 0-12 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHanger0)"/>`),
		g.Group(children),
	)
}