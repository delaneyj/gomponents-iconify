package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.92 1.898C4.744.712 3.727 0 2.484 0A2.47 2.47 0 0 0 .01 2.462a2.468 2.468 0 0 0 2.954 2.414L.975 1.888l3.945.01z"/><path d="m9.966 9.001l6.066-5.975l-13-.026l6.042 6v4.031L7.027 15h5.01l-2.045-1.963l-.026-4.036zM5.188 4l8.594.04l-1.786 1.701c-.47.015-1.142-.087-2.058-.484c-1.219-.442-2.322.405-2.697.743L5.188 4z"/></g>`),
		g.Group(children),
	)
}