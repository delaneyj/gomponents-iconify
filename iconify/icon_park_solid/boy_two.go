package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBoyTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="10" r="6" fill="#fff"/><path d="M31 44v-9l5-3l-4-13s-4-3-8-3s-8 3-8 3l-4 12l5 4v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBoyTwo0)"/>`),
		g.Group(children),
	)
}