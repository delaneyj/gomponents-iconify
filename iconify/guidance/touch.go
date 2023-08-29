package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Touch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9 16v-3.5a3 3 0 0 0-3-3h-.5v11A3.5 3.5 0 0 1 9 24m10.5 0V12.5a3 3 0 0 0-3-3h-.901M16 12v-1a3 3 0 0 0-3-3h-.5v4V.5H12a3 3 0 0 0-3 3V13"/>`),
		g.Group(children),
	)
}