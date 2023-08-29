package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnInsert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30h-6a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v18a2.002 2.002 0 0 1-2 2zm-6-20v18h6V10zm-6-1l5.586-5.586L20.172 2L16 6.172L11.828 2l-1.414 1.414L16 9zm-6 21H4a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v18a2.002 2.002 0 0 1-2 2zM4 10v18h6V10z"/>`),
		g.Group(children),
	)
}