package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flatastic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.71 13.31v21.56a3.63 3.63 0 0 0 3.62 3.63h1.09m-16.59 0V14.57A5.07 5.07 0 0 1 22.9 9.5h0a9.17 9.17 0 0 1 1.64.13m-10.96 9.66h19.93"/><rect width="31" height="39" x="8.5" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}