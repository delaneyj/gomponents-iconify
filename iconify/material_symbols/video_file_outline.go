package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFileOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18h4q.425 0 .713-.288T14 17v-1l2 1.05v-4.1L14 14v-1q0-.425-.288-.712T13 12H9q-.425 0-.713.288T8 13v4q0 .425.288.713T9 18Zm-3 4q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v12q0 .825-.588 1.413T18 22H6Zm7-13V4H6v16h12V9h-5ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}