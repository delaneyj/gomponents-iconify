package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamilyHistoryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.5q-1.575 0-2.663-1.088T8.25 18.75q0-1.3.775-2.288T11 15.125V13H5V9H2.5V2h7v7H7v2h10V8.875q-1.2-.35-1.975-1.338T14.25 5.25q0-1.575 1.088-2.663T18 1.5q1.575 0 2.663 1.088T21.75 5.25q0 1.3-.775 2.288T19 8.874V13h-6v2.125q1.2.35 1.975 1.338t.775 2.287q0 1.575-1.088 2.663T12 22.5ZM18 7q.725 0 1.238-.513t.512-1.237q0-.725-.513-1.238T18 3.5q-.725 0-1.238.513T16.25 5.25q0 .725.513 1.238T18 7ZM4.5 7h3V4h-3v3ZM12 20.5q.725 0 1.238-.513t.512-1.237q0-.725-.513-1.238T12 17q-.725 0-1.238.513t-.512 1.237q0 .725.513 1.238T12 20.5Zm-6-15Zm12-.25Zm-6 13.5Z"/>`),
		g.Group(children),
	)
}