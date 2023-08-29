package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScienceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-1.275 0-1.813-1.137t.263-2.113L9 11V5H8q-.425 0-.713-.288T7 4q0-.425.288-.713T8 3h8q.425 0 .713.288T17 4q0 .425-.288.713T16 5h-1v6l5.55 6.75q.8.975.263 2.113T19 21H5Zm0-2h14l-6-7.3V5h-2v6.7L5 19Zm7-7Z"/>`),
		g.Group(children),
	)
}