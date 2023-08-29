package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Types(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m30 16l-8 8l-1.414-1.414L27.172 16l-6.586-6.586L22 8l8 8z"/><path fill="currentColor" d="M16 22a.997.997 0 0 1-.707-.293l-5-5a1 1 0 0 1 0-1.414l5-5a1 1 0 0 1 1.414 0l5 5a1 1 0 0 1 0 1.414l-5 5A.997.997 0 0 1 16 22Zm-3.586-6L16 19.586L19.586 16L16 12.414Z"/><path fill="currentColor" d="m2 16l8-8l1.414 1.414L4.828 16l6.586 6.586L10 24l-8-8z"/>`),
		g.Group(children),
	)
}