package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDeviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3v4l-3.857 6H18v4h-6.429L9 21H6v-4l3.857-6H6V7h6.429L15 3z"/>`),
		g.Group(children),
	)
}