package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m351 54l-29 28l-81-81l14-1q100 0 173.5 68T510 235h-32q-6-60-40.5-108T351 54zM217 37l257 257q9 9 9 22.5t-9 22.5L338 475q-9 9-22.5 9t-22.5-9L36 218q-9-9-9-22.5t9-22.5L172 37q9-9 22.5-9t22.5 9zm98 415l136-136L195 60L59 196zm-156 6l29-28l81 81l-14 1q-100 0-173.5-68T0 277h32q6 60 40 108t87 73z"/>`),
		g.Group(children),
	)
}