package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M85 207v-95C85 51 140 0 210 0c69 0 125 51 125 112v95H85zm337 6v191c-41 112-162 129-162 129v127h133v57H27v-57h134V533S41 516 0 404V213l49-23v174s15 117 161 117s163-117 163-117V190zm-87 27v86c0 61-56 112-125 112c-70 0-125-51-125-112v-86h250z"/>`),
		g.Group(children),
	)
}