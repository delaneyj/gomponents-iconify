package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sugarizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.642 31.265l6.631-6.631a3.282 3.282 0 0 0-4.642-4.642L24 26.623l-6.631-6.631a3.282 3.282 0 1 0-4.642 4.642l6.631 6.631l-6.631 6.632a3.282 3.282 0 1 0 4.642 4.642L24 35.907l6.631 6.632a3.282 3.282 0 0 0 4.642-4.642Z"/><circle cx="23.999" cy="10.774" r="6.274" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}