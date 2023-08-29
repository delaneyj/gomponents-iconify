package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNotifications(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h2v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q.625.15 1.175.425t1.05.65L10.6 10.4V15h4.6l2.8-2.8V17h2v2H4Zm8 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm.6-9v-1.775l4.975-4.975l1.775 1.775L14.375 13H12.6ZM20 7.375l.85-.85q.15-.15.15-.363t-.15-.362L19.8 4.75q-.15-.15-.363-.15t-.362.15l-.85.85L20 7.375Z"/>`),
		g.Group(children),
	)
}