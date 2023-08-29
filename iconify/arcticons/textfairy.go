package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textfairy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.296 5.5H5.5v3.796m37 0V5.5h-3.796m0 37H42.5v-3.796m-37 0V42.5h3.796m3.189-31.344h23.028M24.16 36.843V11.156"/>`),
		g.Group(children),
	)
}