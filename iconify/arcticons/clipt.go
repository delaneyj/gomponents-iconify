package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.391 43.5h-4.953c-6.96 0-12.603-5.643-12.603-12.603V18.364a3.781 3.781 0 0 1 3.78-3.781h12.687a3.781 3.781 0 0 1 3.78 3.78v11.35a3.781 3.781 0 0 1-3.78 3.78h-2.603a3.781 3.781 0 0 1-3.782-3.78V8.28a3.781 3.781 0 0 1 3.781-3.78h3.864c6.96 0 12.603 5.643 12.603 12.603v13.793"/>`),
		g.Group(children),
	)
}