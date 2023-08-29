package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAddCart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAddCart1" fill="currentColor"><path id="feAddCart2" d="M8 16h8a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V4H5a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1v9h8.775L18 7V5.001c1.145 0 2 .894 2 1.999c0 .146-.017.291-.05.434l-1.151 5c-.21.915-1.052 1.566-2.024 1.566H8.073L8 13.999V16Zm-.5 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm9 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM14 5h1a1 1 0 0 1 0 2h-1v1a1 1 0 0 1-2 0V7h-1a1 1 0 0 1 0-2h1V4a1 1 0 0 1 2 0v1Z"/></g></g>`),
		g.Group(children),
	)
}