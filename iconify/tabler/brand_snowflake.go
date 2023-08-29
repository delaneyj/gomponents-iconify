package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSnowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 21v-5.5l4.5 2.5M10 21v-5.5L5.5 18m-2-3.5L8 12L3.5 9.5m17 0L16 12l4.5 2.5M10 3v5.5L5.5 6M14 3v5.5L18.5 6M12 11l1 1l-1 1l-1-1z"/>`),
		g.Group(children),
	)
}