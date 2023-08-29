package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PewpewTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 31.41L39.655 6.445H11.309l23.532 6.763L42.5 31.41z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.15 30.594l7.997-2.407l.975-7.897l-17.924-5.789l8.952 16.093z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.175 41.555l24.666-1.77l-19.037-4.019L5.5 15.456l4.675 26.099z"/>`),
		g.Group(children),
	)
}