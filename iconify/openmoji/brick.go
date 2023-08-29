package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 45h17v9H15zm18 1h16v8H33zm7-10h17v9H40zm-17 1h16v7H23zm-8-10h17v9H15zm17 0h17v9H32zm8-9h17v9H40zm-17 0h17v9H23z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M32 36v-9m22 0H15m25-8v8m18 9H19m21 9v-9m12 9H19m13 8v-8m15-27H20m3 0v9m0 18v-4m9-5v-9m17 19v5m-5 3H26m31-14v-4"/>`),
		g.Group(children),
	)
}