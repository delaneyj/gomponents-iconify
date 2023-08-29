package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeelgoodOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFeelgoodOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 6H4v30h21l10 5v-5h9V6Z"/><path stroke="#000" d="M13 23s4 4 11 4s11-4 11-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFeelgoodOne0)"/>`),
		g.Group(children),
	)
}