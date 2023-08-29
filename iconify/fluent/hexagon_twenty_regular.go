package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.826 4a.5.5 0 0 0-.433.25l-3.176 5.5a.5.5 0 0 0 0 .5l3.176 5.5a.5.5 0 0 0 .433.25h6.35a.5.5 0 0 0 .433-.25l3.176-5.5a.5.5 0 0 0 0-.5l-3.176-5.5a.5.5 0 0 0-.433-.25h-6.35Zm-1.3-.25a1.5 1.5 0 0 1 1.3-.75h6.35a1.5 1.5 0 0 1 1.3.75l3.175 5.5a1.5 1.5 0 0 1 0 1.5l-3.176 5.5a1.5 1.5 0 0 1-1.299.75h-6.35a1.5 1.5 0 0 1-1.3-.75l-3.175-5.5a1.5 1.5 0 0 1 0-1.5l3.176-5.5Z"/>`),
		g.Group(children),
	)
}