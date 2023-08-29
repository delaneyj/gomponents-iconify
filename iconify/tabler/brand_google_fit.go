package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGoogleFit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9.314L9.657 6.97a3.314 3.314 0 0 0-4.686 4.686L7.314 14L12 18.686l7.03-7.03a3.314 3.314 0 0 0-4.687-4.685L7.313 14"/>`),
		g.Group(children),
	)
}