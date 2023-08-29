package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M203 6q64 52 101 126t37 159q0 70-50 120t-120.5 50T50 411T0 291q0-108 69-190l-1 8q0 33 22.5 56t55.5 23q32 0 52-23t20-56q0-21-3.5-46.5T207 22zm-39 391q43 0 73-30t30-72q0-45-13-86q-30 41-98 55q-29 6-44.5 23.5T96 330q0 28 20 47.5t48 19.5z"/>`),
		g.Group(children),
	)
}