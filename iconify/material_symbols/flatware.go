package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flatware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21V11q-.825 0-1.413-.588T3 9V3.7q0-.3.2-.5t.5-.2q.3 0 .5.2t.2.5V7h.9V3.7q0-.3.2-.5T6 3q.3 0 .5.2t.2.5V7h.9V3.7q0-.3.2-.5t.5-.2q.3 0 .5.2t.2.5V9q0 .825-.588 1.413T7 11v10H5Zm7 0V10.9q-1.05-.5-1.525-1.563T10 7.1q0-1.575.788-2.837T13 3q1.425 0 2.212 1.263T16 7.1q0 1.175-.475 2.238T14 10.9V21h-2Zm5 0V3q1.65 0 2.825 1.175T21 7v6h-2v8h-2Z"/>`),
		g.Group(children),
	)
}