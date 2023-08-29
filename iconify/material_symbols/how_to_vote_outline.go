package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToVoteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20v-4.55l2.75-3.125l1.425 1.425l-2 2.25h13.65l-1.95-2.2l1.425-1.425L21 15.45V20q0 .825-.588 1.413T19 22H5Zm0-2h14v-2H5v2Zm5.625-5.625L7.1 10.85q-.575-.575-.562-1.412t.587-1.413l4.9-4.9q.575-.575 1.425-.6t1.425.55L18.4 6.6q.575.575.6 1.4t-.55 1.4l-5 5q-.575.575-1.412.563t-1.413-.588ZM17 8.025L13.475 4.5l-4.95 4.95l3.525 3.525L17 8.025ZM5 20v-2v2Z"/>`),
		g.Group(children),
	)
}