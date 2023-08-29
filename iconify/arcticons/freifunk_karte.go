package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreifunkKarte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="29.741" cy="24" r="13.759" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.427" cy="28.218" r="8.927" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.967 29.366l-2.986-2.987l2.986-2.986m-2.986 2.987h8.734m-23.666-1.987l1.966 1.966l-1.966 1.965m1.966-1.965H8.373"/>`),
		g.Group(children),
	)
}