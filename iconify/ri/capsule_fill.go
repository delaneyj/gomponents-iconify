package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapsuleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.779 4.222a6 6 0 0 1 0 8.485l-2.122 2.12l-4.95 4.951a6 6 0 0 1-8.485-8.485l7.071-7.071a6 6 0 0 1 8.486 0Zm-4.95 10.606L9.172 9.171l-3.536 3.536a4 4 0 1 0 5.657 5.657l3.536-3.536Z"/>`),
		g.Group(children),
	)
}