package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0h85v43H43v85H0V43zm43 213v85h85v43H43q-18 0-30.5-12.5T0 341v-85h43zm298 85v-85h43v85q0 18-12.5 30.5T341 384h-85v-43h85zm0-341q18 0 30.5 12.5T384 43v85h-43V43h-85V0h85z"/>`),
		g.Group(children),
	)
}