package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.96 7.404L16.596 1.04a.5.5 0 0 0-.707 0L1.04 15.889a.5.5 0 0 0 0 .707l6.364 6.364a.5.5 0 0 0 .707 0l3.18-3.18l.002-.002l2.827-2.827h.001v-.002l2.829-2.827v-.001l2.828-2.828l3.182-3.182a.5.5 0 0 0 0-.707zm-3.535 2.828l-1.768-1.767a.5.5 0 1 0-.707.707l1.768 1.767l-2.122 2.122l-3.182-3.182a.5.5 0 1 0-.707.707l3.182 3.182l-2.121 2.121L12 14.121a.5.5 0 0 0-.707.707l1.767 1.768l-2.12 2.122l-3.183-3.183a.5.5 0 1 0-.707.707l3.182 3.183l-2.475 2.474l-5.656-5.657L16.242 2.101L21.9 7.758l-2.474 2.474z"/>`),
		g.Group(children),
	)
}