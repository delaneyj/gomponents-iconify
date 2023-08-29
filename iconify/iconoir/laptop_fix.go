package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopFix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.8 14.222H3.654a.6.6 0 0 0-.578.437L1.72 19.456A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234Zm0 0v-6.11m-17.6 6.11V4a2 2 0 0 1 2-2H12m-1 17h2m4.657-14.172l-2.829 2.829m5.657-2.829A2 2 0 1 1 17.657 2m-2.829 8.485A2 2 0 0 0 12 7.657"/>`),
		g.Group(children),
	)
}