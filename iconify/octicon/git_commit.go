package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M10.86 7c-.45-1.72-2-3-3.86-3c-1.86 0-3.41 1.28-3.86 3H0v2h3.14c.45 1.72 2 3 3.86 3c1.86 0 3.41-1.28 3.86-3H14V7h-3.14zM7 10.2c-1.22 0-2.2-.98-2.2-2.2c0-1.22.98-2.2 2.2-2.2c1.22 0 2.2.98 2.2 2.2c0 1.22-.98 2.2-2.2 2.2z" fill="currentColor"/>`),
		g.Group(children),
	)
}