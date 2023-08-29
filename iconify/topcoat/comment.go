package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M29.5 30.5h7c2.529 0 3-.529 3-3v-21c0-2.41-.59-3-3-3h-32c-2.47 0-3 .53-3 3v20.971c0 2.469.41 3.029 3 3.029h13s9.562 8.756 10.75 9.812c.422.375 1.281.172 1.25-.812v-9zm-22-9h21v3h-21v-3zm0-6h13v3h-13v-3zm0-6h26v3h-26v-3z"/>`),
		g.Group(children),
	)
}