package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGirlTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="10" r="6" fill="#555"/><path d="M28 44v-8h10L27.23 16h-6.46L10 36h10v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGirlTwo0)"/>`),
		g.Group(children),
	)
}