package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.628.021H2.352a1.33 1.33 0 0 0-1.33 1.331v13.275c0 .734.596 1.331 1.33 1.331h13.276c.734 0 1.33-.597 1.33-1.331V1.352a1.33 1.33 0 0 0-1.33-1.331zm-3.23 13.484l-3.385-1.883l-3.385 1.883l.646-3.987l-2.737-2.824l3.783-.58l1.693-3.628l1.692 3.628l3.784.58l-2.738 2.824l.647 3.987z"/>`),
		g.Group(children),
	)
}