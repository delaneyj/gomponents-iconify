package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a9 9 0 0 1 9 9v7.5a3.5 3.5 0 0 1-6.39 1.976a2.999 2.999 0 0 1-5.223 0a3.5 3.5 0 0 1-6.382-1.783L3 18.499V11a9 9 0 0 1 9-9Zm0 10c-1.105 0-2 1.12-2 2.5s.895 2.5 2 2.5s2-1.12 2-2.5s-.895-2.5-2-2.5ZM9.5 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}