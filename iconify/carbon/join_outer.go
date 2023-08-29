package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinOuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 6a9.954 9.954 0 0 0-4 .838a9.995 9.995 0 0 1 0 18.324A9.999 9.999 0 1 0 20 6zM10 16a9.998 9.998 0 0 1 6-9.162a10 10 0 1 0 0 18.324A9.998 9.998 0 0 1 10 16z"/>`),
		g.Group(children),
	)
}