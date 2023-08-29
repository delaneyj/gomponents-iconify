package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NurseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.957 15.564A8.009 8.009 0 0 1 19.94 22H4.062a8.009 8.009 0 0 1 4.982-6.436L12.001 20l2.956-4.436ZM18.001 2v6A6 6 0 0 1 6 8V2h12Zm-2 6H8a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}