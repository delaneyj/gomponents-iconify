package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SipFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.96 6.501l2.829-2.828a1 1 0 0 1 1.414 0l2.121 2.121a1 1 0 0 1 0 1.414l-2.828 2.829l1.768 1.768l-1.415 1.414l-7.07-7.071l1.413-1.415l1.768 1.768Zm-3.182 2.475l4.243 4.243l-7.778 7.778H3v-4.243l7.778-7.778Z"/>`),
		g.Group(children),
	)
}