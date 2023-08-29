package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mvvapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.55 35.79c1.66.77 3.5 1.21 5.45 1.21c7.18 0 13-5.82 13-13s-5.82-13-13-13s-13 5.82-13 13c0 1.95.44 3.79 1.21 5.45"/>`),
		g.Group(children),
	)
}