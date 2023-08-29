package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineAxisRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 13l-5.25 5.25q-.325.325-.75.325t-.75-.325q-.325-.325-.325-.75t.325-.75L8.8 10.7q.3-.3.7-.3t.7.3l3.3 3.3l1.675-1.9l-5.6-5.175L4.25 12.25q-.325.325-.75.325t-.75-.325q-.325-.325-.325-.75t.325-.75l6.075-6.075Q9.1 4.4 9.5 4.387t.7.263l6.4 5.875l3.3-3.75q.275-.325.713-.325t.737.3q.275.275.288.663t-.263.687l-3.325 3.75l3.15 2.9q.35.3.35.75t-.325.775q-.3.3-.725.3t-.75-.275l-3.1-2.85l-2.45 2.75q-.275.325-.712.338t-.738-.288L9.5 13Z"/>`),
		g.Group(children),
	)
}