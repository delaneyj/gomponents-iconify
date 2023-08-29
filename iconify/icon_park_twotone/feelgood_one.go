package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeelgoodOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFeelgoodOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 6H4v30h21l10 5v-5h9V6Z"/><path d="M13 23s4 4 11 4s11-4 11-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFeelgoodOne0)"/>`),
		g.Group(children),
	)
}