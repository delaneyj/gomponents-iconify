package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonesaver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.17 15.39H9L5.63 31.56h0A10.94 10.94 0 0 0 16.57 42.5H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.63 31.56H12a5 5 0 0 0 5 5h7m0-6.04V5.5M13.08 21.11L24 30.52m5.83-15.13H39l3.41 16.17h0A10.94 10.94 0 0 1 31.43 42.5H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.37 31.56H36a5 5 0 0 1-5 5h-7m10.92-15.45L24 30.52"/>`),
		g.Group(children),
	)
}