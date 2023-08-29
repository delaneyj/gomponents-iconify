package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2042 723l-389 1197H395L6 723L1024 18l1018 705zm-151 51l-867-600l-867 600l331 1018h1072l331-1018z"/>`),
		g.Group(children),
	)
}