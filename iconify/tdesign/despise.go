package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Despise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm5-3h4v2H9v2H7v-2H6V9Zm8 0h4v2h-1v2h-2v-2h-1V9Zm-6.33 6.5A4.998 4.998 0 0 1 12 13a4.998 4.998 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A2.995 2.995 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001l.5-.866Z"/>`),
		g.Group(children),
	)
}