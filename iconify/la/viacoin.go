package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viacoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.281 6l2.563 6H5v2h4.688l.843 2H5v2h6.406l4.469 10.531L20.469 18H27v-2h-5.656l.875-2H27v-2h-3.906l2.625-6H23.53l-4.375 10H12.72L8.438 6zm7.282 12h4.718l-2.375 5.469z"/>`),
		g.Group(children),
	)
}