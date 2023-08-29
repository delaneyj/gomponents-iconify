package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RvHookupOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.5 21.5l-1.4-1.4l1.1-1.1h-7.4q-.3.9-1.075 1.45T9 21q-.95 0-1.712-.55T6.2 19H2v-9h6V7H2V5h15v12h2.15l-1.05-1.05l1.4-1.45L23 18l-3.5 3.5ZM10 10h5V7h-5v3Zm-1 9q.425 0 .713-.288T10 18q0-.425-.288-.713T9 17q-.425 0-.713.288T8 18q0 .425.288.713T9 19Zm-2.8-2q.325-.9 1.087-1.45T9 15q.95 0 1.725.55T11.8 17H15v-5H4v5h2.2Zm0-5H4h11h-8.8Z"/>`),
		g.Group(children),
	)
}