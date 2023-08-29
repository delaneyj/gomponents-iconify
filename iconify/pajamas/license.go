package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func License(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 10.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM14 6a5.997 5.997 0 0 1-2.886 5.13l.58 3.185L12 16l-1.623-.544L8 14.66l-2.377.796L4 16l.306-1.684l.58-3.187A6 6 0 1 1 14 6Zm-7.748 6h3.496l.322 1.772l-1.594-.534l-.476-.16l-.476.16l-1.594.534L6.252 12ZM9.5 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM11 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}