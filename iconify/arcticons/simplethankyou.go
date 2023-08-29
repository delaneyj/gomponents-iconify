package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplethankyou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 14.81C-.49 26.4 22.37 39.47 24 39.47c1.79 0 24.48-13.07 18.5-24.66c-2.63-4.46-9.5-11.56-18.5 0c-5.38-7.81-13.73-8.42-18.5 0Z"/>`),
		g.Group(children),
	)
}