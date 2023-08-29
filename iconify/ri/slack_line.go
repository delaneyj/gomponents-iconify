package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlackLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.501 3a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-3 0v-5a1.5 1.5 0 0 1 1.5-1.5Zm-10 10h1.5v1.5a1.5 1.5 0 1 1-1.5-1.5Zm8.5 5h1.5a1.5 1.5 0 1 1-1.5 1.5V18Zm1.5-5h5a1.5 1.5 0 0 1 0 3h-5a1.5 1.5 0 1 1 0-3Zm5-5a1.5 1.5 0 0 1 0 3h-1.5V9.5a1.5 1.5 0 0 1 1.5-1.5Zm-15 0h5a1.5 1.5 0 1 1 0 3h-5a1.5 1.5 0 0 1 0-3Zm5-5a1.5 1.5 0 0 1 1.5 1.5V6h-1.5a1.5 1.5 0 1 1 0-3Zm0 10a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-3 0v-5a1.5 1.5 0 0 1 1.5-1.5Z"/>`),
		g.Group(children),
	)
}