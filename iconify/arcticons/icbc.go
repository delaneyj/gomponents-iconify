package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.27 42.5H5.5V28.44h12.46v-8.88H5.5V5.5h15.77m5.46 0H42.5v14.06H30.04v8.88H42.5V42.5H26.73"/>`),
		g.Group(children),
	)
}