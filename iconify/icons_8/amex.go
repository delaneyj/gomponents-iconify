package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3H5zm0 2h22c.565 0 1 .435 1 1v14c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1zm2.063 5.25L5 18.75h1.25l.406-1.25H9l.406 1.25h2.344v-4.125l1.5 4.125h1.094l1.53-4v4h1.095v-5.5h-1.783l-1.376 3.72l-1.373-3.72H10.5v5.22l-1.938-5.22h-1.5zm11.406 0v5.5h4.404l1.375-1.78l1.374 1.78H27L24.937 16L27 13.25h-1.5l-1.375 1.656L23 13.25h-4.53zM7.75 14.344l.688 1.937H7.062l.688-1.935zm11.813.156h2.75l1.125 1.5l-1.25 1.656h-2.625v-1.093h2.5v-1.125h-2.5V14.5z"/>`),
		g.Group(children),
	)
}