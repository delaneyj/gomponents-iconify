package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.1 21.875l-6.375-6.35L7.55 20.7q-.15.15-.337.225T6.825 21H4q-.425 0-.713-.288T3 20v-2.825q0-.2.075-.388t.225-.337l5.175-5.175L2.1 4.9q-.275-.275-.275-.7t.3-.725q.275-.275.7-.275t.725.275l16.975 17q.275.275.275.7t-.275.7q-.3.3-.725.3t-.7-.3Zm.2-12.95l-4.25-4.2l1.4-1.4q.575-.575 1.413-.575t1.412.575l1.4 1.4q.575.575.6 1.388t-.55 1.387L19.3 8.925Zm-3.725 3.75l-4.25-4.25L13.6 6.15l4.25 4.25l-2.275 2.275Z"/>`),
		g.Group(children),
	)
}