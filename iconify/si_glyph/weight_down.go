package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.905 5.01c0-1.088-.912-1.971-2.038-1.971h-.923c.034-.152.056-.32.056-.508C10 1.135 8.877 0 7.495 0S4.99 1.135 4.99 2.531c0 .188.026.355.068.508h-.922c-1.125 0-2.037.883-2.037 1.971L.083 13.985c0 1.09.912 1.972 2.039 1.972H12.88c1.126 0 2.037-.882 2.037-1.972L12.905 5.01zM5.824 2.531c0-.949.749-1.723 1.671-1.723c.921 0 1.67.773 1.67 1.723c0 .178-.034.346-.083.508H5.906a1.743 1.743 0 0 1-.082-.508zm1.139 7.416V6.958h1.074v2.968l2.07.021l-2.648 3.092l-2.49-3.092h1.994z"/>`),
		g.Group(children),
	)
}