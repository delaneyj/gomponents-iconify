package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q-3.525-2.6-5.262-5.05T5 9.15q0-1.775.638-3.113T7.274 3.8q1-.9 2.25-1.35T12 2q1.225 0 2.475.45t2.25 1.35q1 .9 1.638 2.238T19 9.15q0 2.35-1.738 4.8T12 19Zm0-8q.825 0 1.413-.588T14 9q0-.825-.588-1.413T12 7q-.825 0-1.413.588T10 9q0 .825.588 1.413T12 11ZM5 22v-2h14v2H5Z"/>`),
		g.Group(children),
	)
}