package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m249.17 117.7l-83.716 108.032H0V1082.3h1200V225.732H591.65L507.935 117.7H249.17z"/>`),
		g.Group(children),
	)
}