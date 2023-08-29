package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.293 38.023C37.14 38.023 43.5 31.745 43.5 24S37.14 9.977 29.293 9.977S15.085 16.255 15.085 24s6.36 14.023 14.207 14.023Zm-14.15-3.519c5.879 0 10.644-4.703 10.644-10.505s-4.765-10.503-10.644-10.503S4.5 18.199 4.5 23.999s4.765 10.505 10.643 10.505Z"/>`),
		g.Group(children),
	)
}