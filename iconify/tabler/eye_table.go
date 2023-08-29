package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 18h-.011M12 18h-.011M16 18h-.011M4 3h16M5 3v17a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V3m-5 4h-4m-1 8h1m4 0h1m-3-4V7"/>`),
		g.Group(children),
	)
}