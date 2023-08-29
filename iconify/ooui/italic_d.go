package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItalicD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 19L6 1h4.6c2.29 0 4.17.59 5.46 1.76C17.36 3.94 18 5.72 18 8.09c0 2.14-.42 4.03-1.28 5.68a9.36 9.36 0 0 1-3.74 3.85A11.98 11.98 0 0 1 7 19H2Zm5.16-2.3c1.71 0 3.16-.38 4.33-1.15a7.37 7.37 0 0 0 2.66-3.1c.6-1.3.9-2.74.9-4.34c0-1.68-.43-2.9-1.28-3.66a4.86 4.86 0 0 0-3.4-1.16H8.33l-3 13.42h1.83Z"/>`),
		g.Group(children),
	)
}