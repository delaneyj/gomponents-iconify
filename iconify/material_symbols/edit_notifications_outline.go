package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNotificationsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-8-3v-2h2v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q.625.15 1.175.425t1.05.65l-1.45 1.45q-.5-.35-1.062-.537T12 6q-1.65 0-2.825 1.175T8 10v7h8v-2.8l2-2V17h2v2H4Zm8-7.5Zm.6 1.5v-1.775l4.975-4.975l1.775 1.775L14.375 13H12.6ZM20 7.375l.85-.85q.15-.15.15-.363t-.15-.362L19.8 4.75q-.15-.15-.363-.15t-.362.15l-.85.85L20 7.375Z"/>`),
		g.Group(children),
	)
}