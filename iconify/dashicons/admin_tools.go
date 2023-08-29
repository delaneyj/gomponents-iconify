package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.68 9.77a4.543 4.543 0 0 1-4.95.99l-5.41 6.52c-.99.99-2.59.99-3.58 0s-.99-2.59 0-3.57l6.52-5.42c-.68-1.65-.35-3.61.99-4.95c1.28-1.28 3.12-1.62 4.72-1.06l-2.89 2.89l2.82 2.82l2.86-2.87c.53 1.58.18 3.39-1.08 4.65zM3.81 16.21c.4.39 1.04.39 1.43 0c.4-.4.4-1.04 0-1.43c-.39-.4-1.03-.4-1.43 0a1.02 1.02 0 0 0 0 1.43z"/>`),
		g.Group(children),
	)
}