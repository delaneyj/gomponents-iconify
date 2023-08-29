package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peercoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.236 43.5C5.134 41.333 6.634 26.837 7.162 4.514c32.472-.768 43.31 29.18 25.128 37.766c3-15.783-4.366-26.461-16.615-32.386c8.816 6.701 17.047 18.781 11.56 33.606Z"/>`),
		g.Group(children),
	)
}