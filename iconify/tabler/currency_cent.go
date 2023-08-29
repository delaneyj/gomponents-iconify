package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyCent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.007 7.54A5.965 5.965 0 0 0 11.999 6a6 6 0 0 0-5.992 6c0 3.314 2.682 6 5.992 6a5.965 5.965 0 0 0 4-1.536M12 20v-2m0-12V4"/>`),
		g.Group(children),
	)
}