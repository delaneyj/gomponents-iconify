package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeFillSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.97 13.5l4.705 4.626L19.61 30.27l-6.098 1.34l1.682-6.071L26.97 13.5ZM14.064 34.187c8.278 1.07 13.177-.646 15.272-4.443c.694-1.257-.76-3.033-2.47-.998c-2.277 2.707-.653 9.324 5.204 2.523c-1.156 3.473-.139 3.585 2.418 1.314"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/>`),
		g.Group(children),
	)
}