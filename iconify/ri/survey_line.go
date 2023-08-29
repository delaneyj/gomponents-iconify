package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2v2h3.007c.548 0 .993.445.993.993v16.014a.994.994 0 0 1-.993.993H3.993A.993.993 0 0 1 3 21.007V4.993C3 4.445 3.445 4 3.993 4H7V2h10ZM7 6H5v14h14V6h-2v2H7V6Zm2 10v2H7v-2h2Zm0-3v2H7v-2h2Zm0-3v2H7v-2h2Zm6-6H9v2h6V4Z"/>`),
		g.Group(children),
	)
}