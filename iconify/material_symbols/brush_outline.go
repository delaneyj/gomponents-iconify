package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-1.125 0-2.225-.55T2 19q.65 0 1.325-.513T4 17q0-1.25.875-2.125T7 14q1.25 0 2.125.875T10 17q0 1.65-1.175 2.825T6 21Zm0-2q.825 0 1.413-.588T8 17q0-.425-.288-.713T7 16q-.425 0-.713.288T6 17q0 .575-.138 1.05t-.362.9q.125.05.25.05H6Zm5.75-4L9 12.25l8.95-8.95q.275-.275.688-.288t.712.288l1.35 1.35q.3.3.3.7t-.3.7L11.75 15ZM7 17Z"/>`),
		g.Group(children),
	)
}