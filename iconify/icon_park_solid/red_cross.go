package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRedCross0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#fff" stroke="#fff"/><path fill="#000" stroke="#000" d="M27 12h-6v9h-9v6h9v9h6v-9h9v-6h-9v-9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRedCross0)"/>`),
		g.Group(children),
	)
}