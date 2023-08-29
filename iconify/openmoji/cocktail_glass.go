package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CocktailGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M45.686 20.066h7.982m-35.624 0h19.042"/><path fill="#B1CC33" d="m17.061 19l6.439 8.5L36 44l12.5-16.5l6.439-8.5z"/><path fill="#5C9E31" d="M44.939 19L38.5 27.5L31 37.4l5 6.6l12.5-16.5l6.439-8.5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M36 44v23m-16 0h32m9-56L48.5 27.5L36 44L23.5 27.5L11.239 11.315m0 0L11 11m28 12L55 4"/>`),
		g.Group(children),
	)
}