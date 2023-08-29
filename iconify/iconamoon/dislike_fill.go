package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.277 19.219A4 4 0 0 0 14.606 21h.213a2 2 0 0 0 1.973-2.329L16.18 15h2.38a3 3 0 0 0 2.942-3.588l-1.2-6A3 3 0 0 0 17.36 3H6a3 3 0 0 0-3 3v8a1 1 0 0 0 1 1h3.93a1 1 0 0 1 .832.445l2.515 3.774ZM7 5v8H5V6a1 1 0 0 1 1-1h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}