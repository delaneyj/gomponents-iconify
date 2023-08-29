package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifiautomatic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.25 23.51a14.44 14.44 0 0 0-18.45-.05v.05m9.2 3.54a7.61 7.61 0 1 0 7.61 7.61h0A7.61 7.61 0 0 0 24 27.05Zm14-9.36a22 22 0 0 0-28 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.48 12.4a29 29 0 0 0-37 0m15.23 26.47l3.39-8.44m3.24 8.47l-3.24-8.47m2.16 5.63h-4.43"/>`),
		g.Group(children),
	)
}