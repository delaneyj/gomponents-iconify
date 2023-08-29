package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.097 12c1.441 4.08 5.332 7 9.903 7c4.57 0 8.46-2.92 9.902-7C20.461 7.92 16.57 5 12 5c-4.57 0-8.462 2.92-9.903 7Zm-2.008-.304C1.7 6.654 6.421 3 12 3c5.58 0 10.302 3.654 11.91 8.696l.098.304l-.097.304C22.3 17.346 17.578 21 12 21C6.42 21 1.698 17.346.09 12.304L-.009 12l.097-.304ZM12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/>`),
		g.Group(children),
	)
}