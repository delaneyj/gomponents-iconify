package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNotificationsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h1v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q.625.15 1.175.425t1.05.65L11.05 9.95q-.2.2-.325.488t-.125.587V13.5q0 .625.438 1.063T12.1 15h2.475q.3 0 .588-.125t.487-.325L18 12.2V17h1q.425 0 .713.288T20 18q0 .425-.288.713T19 19H5Zm7 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm1.1-9q-.2 0-.35-.15t-.15-.35v-.85q0-.2.075-.388t.225-.337l4.675-4.675l1.775 1.775l-4.675 4.675q-.15.15-.338.225T13.95 13h-.85ZM20 7.375l.85-.85q.15-.15.15-.363t-.15-.362L19.8 4.75q-.15-.15-.363-.15t-.362.15l-.85.85L20 7.375Z"/>`),
		g.Group(children),
	)
}