package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 341V21h213v320H107zM0 299V64h85v235H0zM341 64h86v235h-86V64z"/>`),
		g.Group(children),
	)
}