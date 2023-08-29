package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplecamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.63 15.22l-24.73-.05M13.8 24L26.15 2.61M6.53 11.44l12.34 21.34m-14.5 0l24.71.05m-7.29 12.55l12.4-21.36m7.24 12.58L29.1 15.17"/>`),
		g.Group(children),
	)
}