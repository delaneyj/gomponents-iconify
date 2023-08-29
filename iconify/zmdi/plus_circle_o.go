package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M235 109v86h85v42h-85v86h-43v-86h-85v-42h85v-86h43zM213 3q88 0 151 62.5T427 216t-63 150.5T213 429T62.5 366.5T0 216T62.5 65.5T213 3zm.5 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50z"/>`),
		g.Group(children),
	)
}