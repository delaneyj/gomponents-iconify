package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneTiltFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.4 44.34v.15l-58.2 191.94a15.88 15.88 0 0 1-14 11.51q-.69.06-1.38.06a15.86 15.86 0 0 1-14.42-9.15l-35.71-75.39a4 4 0 0 1 .79-4.54l57.26-57.27a8 8 0 0 0-11.31-11.31L97.08 147.6a4 4 0 0 1-4.54.79l-75-35.53A16.37 16.37 0 0 1 8 97.36a15.89 15.89 0 0 1 11.57-14.52l191.94-58.2h.15a16 16 0 0 1 19.74 19.7Z"/>`),
		g.Group(children),
	)
}