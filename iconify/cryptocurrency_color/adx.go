package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#1B75BC"/><path fill="#FFF" d="M11.587 8.602L16 13.208l4.305-4.606L23 11.36L18.604 16L23 20.603l-2.695 2.793L16 18.792l-4.413 4.604L9 20.68l4.36-4.706L9 11.36l2.587-2.757zm.355-.376L15.982 4l4.04 4.226l-2.126 2.265l-1.914-2.114l-1.914 2.114l-2.126-2.265zm0 15.548l2.126-2.265l1.914 2.114l1.914-2.114l2.126 2.265L15.982 28l-4.04-4.226z"/></g>`),
		g.Group(children),
	)
}