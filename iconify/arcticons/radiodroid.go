package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiodroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.77 7.16a16.84 16.84 0 0 1 0 33.68m0-5.62a11.22 11.22 0 1 0 0-22.44M20 7.51c.77-.11 1.16.71 1.16 2.46V38c0 2.81-1 3.21-3 1.24L5.72 26.81a4.09 4.09 0 0 1 0-5.62L18.18 8.73A3.89 3.89 0 0 1 20 7.51Zm6.78 10.56a5.93 5.93 0 0 1 0 11.86"/>`),
		g.Group(children),
	)
}