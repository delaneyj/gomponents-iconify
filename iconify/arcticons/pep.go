package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2a22 22 0 1 0 22 22A22 22 0 0 0 24 2Zm-2.88 27.78h5.76m-5.76-11.59h5.76m-5.76 5.79h5.76"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.93 29.75V18.16h3.86a3.89 3.89 0 0 1 0 7.78H9.93m20.41 3.81V18.16h3.87a3.89 3.89 0 0 1 0 7.78h-3.87"/>`),
		g.Group(children),
	)
}