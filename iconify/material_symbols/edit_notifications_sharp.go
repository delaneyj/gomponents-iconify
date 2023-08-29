package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNotificationsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h2v-7q0-2.075 1.25-3.688T10.5 4.2V2h3v2.2q.625.15 1.175.425t1.05.65L10.6 10.4V15h4.6l2.8-2.8V17h2v2H4Zm8 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm.6-9v-1.775l4.975-4.975l1.775 1.775L14.375 13H12.6ZM20 7.375l1.225-1.225l-1.775-1.775L18.225 5.6L20 7.375Z"/>`),
		g.Group(children),
	)
}