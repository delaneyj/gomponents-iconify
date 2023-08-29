package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19h3v-5q0-.425.288-.713T10 13h4q.425 0 .713.288T15 14v5h3v-9l-6-4.5L6 10v9Zm-2 0v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21h-4q-.425 0-.713-.288T13 20v-5h-2v5q0 .425-.288.713T10 21H6q-.825 0-1.413-.588T4 19Zm8-6.75Z"/>`),
		g.Group(children),
	)
}