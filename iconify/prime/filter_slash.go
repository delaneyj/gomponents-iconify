package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.17 3.91a.76.76 0 0 0-.67-.41h-7a.75.75 0 0 0 0 1.5H18l-1.85 2.56a.74.74 0 1 0 1.2.88l2.75-3.75a.73.73 0 0 0 .07-.78ZM8.65 3.72H8.6a.76.76 0 0 0-.19-.12a.79.79 0 0 0-.22-.05H4.5a.76.76 0 0 0-.67.41a.73.73 0 0 0 .07.78L9.25 12v7.75a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V12l.06-.09l3.66 3.62a.74.74 0 0 0 .53.22a.73.73 0 0 0 .53-.22a.75.75 0 0 0 0-1.06Zm4.75 7.59a.71.71 0 0 0-.15.44V19h-2.5v-7.25a.71.71 0 0 0-.15-.44L6 5h1.82l5.91 5.85Z"/>`),
		g.Group(children),
	)
}