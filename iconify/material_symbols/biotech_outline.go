package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BiotechOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21v-2h5v-2q-2.075 0-3.538-1.463T5 12q0-1.525.838-2.775T8.1 7.4q.2-.85.888-1.375T10.55 5.5L10 3.95l.95-.35l-.35-.9l1.9-.7l.3.95l.95-.35l2.75 7.5l-.95.35l.35.95l-1.9.7l-.3-.95l-.95.35l-.6-1.65q-.375.35-.863.525t-.987.125q-.55-.05-1.025-.338T8.45 9.45q-.675.4-1.062 1.075T7 12q0 1.25.875 2.125T10 15h8v2h-5v2h6v2H5Zm8.65-11.45l.9-.35l-1.7-4.7l-.95.35l1.75 4.7ZM10.5 9q.425 0 .713-.288T11.5 8q0-.425-.288-.713T10.5 7q-.425 0-.713.288T9.5 8q0 .425.288.713T10.5 9Zm3.15.55ZM10.5 8Zm0 0Z"/>`),
		g.Group(children),
	)
}