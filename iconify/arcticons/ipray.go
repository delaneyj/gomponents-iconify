package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.234 39.645L24 42.5v-37l18.234 5.481v28.664zM24 11.781l18.233 5.139M24 16.806l18.233 4.225M5.766 39.645L24 42.5v-37L5.766 10.981v28.664zM24 11.781L5.766 16.92M24 16.806L5.766 21.031"/>`),
		g.Group(children),
	)
}