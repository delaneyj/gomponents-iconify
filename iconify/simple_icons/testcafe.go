package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Testcafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.315 4.319l-8.69 8.719l-3.31-3.322l-2.069 2.076l5.379 5.398l10.76-10.796zM5.849 14.689L0 19.682h24l-5.864-4.991h-3.2l-1.024.896h3.584l3.072 2.815H3.417l3.072-2.815h2.688l-.896-.896z"/>`),
		g.Group(children),
	)
}