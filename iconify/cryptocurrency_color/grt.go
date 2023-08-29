package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><filter id="cryptocurrencyColorGrt0" color-interpolation-filters="auto"><feColorMatrix in="SourceGraphic" values="0 0 0 0 1.000000 0 0 0 0 1.000000 0 0 0 0 1.000000 0 0 0 1.000000 0"/></filter></defs><g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#5942CC" fill-rule="nonzero"/><g filter="url(#cryptocurrencyColorGrt0)"><path fill="#000" fill-rule="nonzero" d="M20.707 19.543a1 1 0 0 1 0 1.414l-4 4a1 1 0 1 1-1.414-1.414l4-4a1 1 0 0 1 1.414 0zM15 7.25a6 6 0 1 1 0 12a6 6 0 0 1 0-12zm0 2a4 4 0 1 0 0 8a4 4 0 0 0 0-8zm7-2a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/></g></g>`),
		g.Group(children),
	)
}