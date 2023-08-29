package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 13a1 1 0 1 0 0 2h5a1 1 0 0 0 1-1V9a1 1 0 1 0-2 0v2.586l-4.293-4.293a1 1 0 0 0-1.414 0L8 9.586L3.707 5.293a1 1 0 0 0-1.414 1.414l5 5a1 1 0 0 0 1.414 0L11 9.414L14.586 13H12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}