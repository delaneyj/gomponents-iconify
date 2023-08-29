package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 6v8.5h-9v-13H8v2.75C8 5.216 8.784 6 9.75 6h2.75Zm-.121-1.5L9.5 1.621V4.25c0 .138.112.25.25.25h2.629ZM2 1a1 1 0 0 1 1-1h6.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V15a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}