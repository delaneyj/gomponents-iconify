package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GavelRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.975 21q-.425 0-.7-.288T4 20q0-.425.288-.713T5 19h10.025q.425 0 .7.288T16 20q0 .425-.288.713T15 21H4.975Zm3.25-6.275L5.4 11.9q-.575-.575-.588-1.413t.563-1.412L6.1 8.35L11.8 14l-.725.725q-.575.575-1.425.575t-1.425-.575ZM16 9.8l-5.65-5.7l.725-.725q.575-.575 1.413-.562T13.9 3.4l2.825 2.825q.575.575.575 1.425t-.575 1.425L16 9.8Zm3.9 9.5L7.55 6.95l1.4-1.4l12.375 12.375q.275.275.263.687t-.288.688q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}