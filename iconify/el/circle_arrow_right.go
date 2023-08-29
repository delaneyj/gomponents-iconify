package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.635 0 0 268.635 0 600s268.635 600 600 600s600-268.635 600-600S931.365 0 600 0zm-66.647 297.145L987.667 600L533.353 902.855V739.678h-321.02v-279.34h321.02V297.145z"/>`),
		g.Group(children),
	)
}