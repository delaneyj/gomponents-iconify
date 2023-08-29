package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-8-3v-2h2v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q.45.125.875.288t.8.412q-.7.875-1.063 1.925T13.75 9q0 2.05 1.163 3.725t3.087 2.4V17h2v2H4Zm15-6v-3h-3V8h3V5h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}