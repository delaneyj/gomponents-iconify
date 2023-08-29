package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pressthis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.76 1a3.246 3.246 0 0 1 0 6.49c-.23 0-.47-.03-.7-.08L13 8.47V19H2V4h9.54c.13-2 1.52-3 3.22-3zm0 5.49a2.245 2.245 0 0 0 0-4.49c-1.24 0-2.24 1.01-2.24 2.25c0 .37.1.72.27 1.03L9.57 8.5c-.28.28-1.77 2.22-1.5 2.49c.02.03.06.04.1.04c.49 0 2.14-1.28 2.39-1.53l3.24-3.24c.29.14.61.23.96.23z"/>`),
		g.Group(children),
	)
}