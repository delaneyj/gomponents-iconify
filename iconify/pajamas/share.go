package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 1.5a2.5 2.5 0 1 0-2.469-2.104L5.298 6.263a2.5 2.5 0 1 0 0 3.475l4.733 2.366a2.5 2.5 0 1 0 .671-1.341L5.97 8.396a2.519 2.519 0 0 0 0-.792l4.733-2.367c.455.47 1.092.763 1.798.763Zm1 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM4.5 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}