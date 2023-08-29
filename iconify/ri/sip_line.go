package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.457 18.954l8.564-8.564l-1.414-1.414l-8.564 8.564l1.414 1.414Zm5.735-11.392l-1.414-1.414l1.414-1.415l1.768 1.768l2.829-2.828a1 1 0 0 1 1.414 0l2.121 2.121a1 1 0 0 1 0 1.414l-2.828 2.829l1.768 1.768l-1.415 1.414l-1.414-1.414l-9.192 9.192H3v-4.243l9.192-9.192Z"/>`),
		g.Group(children),
	)
}