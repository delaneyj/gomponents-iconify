package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightContractSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.293 8.5l-.647.647a.5.5 0 1 0 .708.707l1.5-1.5a.5.5 0 0 0 0-.707l-1.5-1.5a.5.5 0 1 0-.708.707l.647.646H4.502a.5.5 0 1 0 0 1h1.79ZM12 13.001a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v6.002a2 2 0 0 0 2 2h8Zm1-2a1 1 0 0 1-1 1H9.998V4H12a1 1 0 0 1 1 1v6.002ZM8.998 4v8.002H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h4.998Z"/>`),
		g.Group(children),
	)
}