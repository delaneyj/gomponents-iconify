package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ht(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#2A3069"/><g fill="#FFF"><path d="M18.347 10.533c0-3.52-1.707-6.506-2.987-7.466c0 0-.107-.107-.107.106c-.106 6.72-3.52 8.534-5.333 11.094c-4.373 5.653-.32 11.946 3.84 13.12c2.347.64-.533-1.174-.853-4.907c-.534-4.693 5.44-8.107 5.44-11.947"/><path d="M20.587 12.88c-.214.853-1.067 2.667-2.24 4.373c-3.947 5.654-1.707 8.427-.427 9.92c.747.854 0 0 1.813-.853c.107-.107 3.627-1.92 3.947-6.08c.427-4.053-2.133-6.613-3.093-7.36"/></g></g>`),
		g.Group(children),
	)
}