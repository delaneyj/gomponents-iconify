package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 11.617a6.818 6.818 0 0 1-6.813 6.817H6.816a6.817 6.817 0 1 1 0-13.634h10.366a6.818 6.818 0 0 1 6.817 6.813zm-6.817-4.545a4.542 4.542 0 1 0 0 9.084a4.542 4.542 0 0 0 .002-9.084h-.002z"/>`),
		g.Group(children),
	)
}