package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10.086L7.207 5.293L5.793 6.707L12 12.914l6.207-6.207l-1.414-1.414L12 10.086ZM18 17H6v-2h12v2Z"/>`),
		g.Group(children),
	)
}