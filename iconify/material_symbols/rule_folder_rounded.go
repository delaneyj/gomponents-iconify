package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleFolderRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm3.825-6.45l-.725-.725q-.3-.3-.7-.287t-.7.312q-.275.3-.288.7t.288.7l1.425 1.425q.3.3.7.3t.7-.3l3.55-3.55q.3-.3.3-.7t-.3-.7q-.3-.3-.712-.3t-.713.3L7.825 13.55ZM16 11.975l-.9-.9q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7l.9.9l-.9.9q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l.9-.9l.9.9q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-.9-.9l.9-.9q.275-.275.275-.7t-.275-.7q-.275-.275-.7-.275t-.7.275l-.9.9Z"/>`),
		g.Group(children),
	)
}