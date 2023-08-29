package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M587 0L478 98s36 22 69 53c32 31 55 66 55 66l93-111s-1-20-45-63C605 0 587 0 587 0zM232 549l-5 48l346-360s-23-32-51-59c-29-28-62-50-62-50L114 488l54-3l6 58zM0 702l86-176l48-4l4 55l58 8l-6 44l-174 88z"/>`),
		g.Group(children),
	)
}