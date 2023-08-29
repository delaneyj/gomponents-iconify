package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActionCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h2048v1664h-640l-384 384l-384-384H0V0zm1920 1536V128H128v1408h565l331 331l331-331h565z"/>`),
		g.Group(children),
	)
}