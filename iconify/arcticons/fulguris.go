package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fulguris(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.32 4.5L20.167 5.98l-5.355 5.602l5.38 5.7l-5.684 5.783l3.996 5.775L12.68 43.5l13.458-12.708l-2.077-5.656l7.759-5.866l-3.749-5.758L35.32 4.5Z"/>`),
		g.Group(children),
	)
}