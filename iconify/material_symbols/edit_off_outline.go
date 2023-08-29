package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.1 21.875l-6.375-6.35L7.25 21H3v-4.25l5.475-5.475l-6.35-6.35L3.55 3.5l16.975 16.975l-1.425 1.4ZM10.6 13.4l-.7-.7l.7.7l.7.7l-.7-.7Zm4.975-.725L14.15 11.25l.875-.875l-1.4-1.4l-.875.875l-1.425-1.425L13.6 6.15l4.25 4.25l-2.275 2.275Zm3.725-3.75l-4.25-4.2l1.4-1.4q.575-.575 1.413-.575t1.412.575l1.4 1.4q.575.575.6 1.388t-.55 1.387L19.3 8.925Zm-5.85 1.625ZM5 19h1.4l4.9-4.9l-1.4-1.4L5 17.6V19Z"/>`),
		g.Group(children),
	)
}