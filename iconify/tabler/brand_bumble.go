package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBumble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12h10M9 8h6m-5 8h4m2.268-13H7.732a1.46 1.46 0 0 0-1.268.748l-4.268 7.509a1.507 1.507 0 0 0 0 1.486l4.268 7.509c.26.462.744.747 1.268.748h8.536a1.46 1.46 0 0 0 1.268-.748l4.268-7.509a1.507 1.507 0 0 0 0-1.486l-4.268-7.509A1.46 1.46 0 0 0 16.268 3z"/>`),
		g.Group(children),
	)
}