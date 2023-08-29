package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5.97 3l-.314.344S4.9 4 4.406 4v2c.68 0 1.15-.276 1.594-.53V10h2V3H5.97zM11 6v2h17V6H11zm-4.5 6A2.497 2.497 0 0 0 4 14.5v.5h2v-.5c0-.217.283-.5.5-.5c.217 0 .5.283.5.5l-.03.03l-.064.064l-2.593 2.5l-.313.28V19h5v-2H7.28l.876-.875l.125-.094l-.03-.03c.502-.41.75-1.02.75-1.5A2.499 2.499 0 0 0 6.5 12zm4.5 3v2h17v-2H11zm-7 6v2h1.375l-.25.406l-.125.22V25h1.5c.217 0 .5.283.5.5c0 .217-.283.5-.5.5H4v2h2.5C7.883 28 9 26.883 9 25.5c0-1.005-.678-1.696-1.53-2.094l.405-.687l.125-.25V21H4zm7 3v2h17v-2H11z"/>`),
		g.Group(children),
	)
}