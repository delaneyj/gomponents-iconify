package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.1 21.875l-6.375-6.35L7.55 20.7q-.15.15-.337.225T6.825 21H4q-.425 0-.713-.288T3 20v-2.825q0-.2.075-.388t.225-.337l5.175-5.175L2.1 4.9q-.3-.3-.288-.7t.313-.7q.3-.3.713-.3t.712.3l16.975 16.975q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM10.6 13.4l-.7-.7l.7.7l.7.7l-.7-.7Zm4.975-.725L14.15 11.25l.875-.875l-1.4-1.4l-.875.875l-1.425-1.425L13.6 6.15l4.25 4.25l-2.275 2.275Zm3.725-3.75l-4.25-4.2l1.4-1.4q.575-.575 1.413-.575t1.412.575l1.4 1.4q.575.575.6 1.388t-.55 1.387L19.3 8.925Zm-5.85 1.625ZM5 19h1.4l4.9-4.9l-1.4-1.4L5 17.6V19Z"/>`),
		g.Group(children),
	)
}