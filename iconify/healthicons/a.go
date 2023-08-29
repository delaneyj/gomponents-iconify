package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func A(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 10a2 2 0 0 1 1.846 1.23l7.483 17.96l.035.084l2.482 5.957a2 2 0 0 1-3.692 1.538L30.167 32H17.833l-1.987 4.77a2 2 0 0 1-3.692-1.54l2.482-5.956a2.01 2.01 0 0 1 .035-.085l7.483-17.958A2 2 0 0 1 24 10Zm-4.5 18h9L24 17.2L19.5 28Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}