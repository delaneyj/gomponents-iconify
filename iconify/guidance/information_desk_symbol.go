package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationDeskSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 7V5M9 9.5h3v8m0 0H8m4 0h4m-4 6C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		g.Group(children),
	)
}