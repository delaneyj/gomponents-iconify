package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseGalleryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.8 16.2l1.4-1.4l-3.2-3.2V7H8v5.4l3.8 3.8Zm6.2 4.3v-2.2q1.85-.875 2.925-2.575T22 12q0-2.025-1.075-3.725T18 5.7V3.5q2.725.95 4.362 3.288T24 12q0 2.875-1.638 5.213T18 20.5ZM9 21q-1.875 0-3.513-.713t-2.85-1.924Q1.425 17.15.712 15.513T0 12q0-1.875.713-3.513t1.925-2.85q1.212-1.212 2.85-1.924T9 3q1.875 0 3.513.713t2.85 1.924q1.212 1.213 1.924 2.85T18 12q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T9 21Zm0-2q2.925 0 4.963-2.038T16 12q0-2.925-2.038-4.963T9 5Q6.075 5 4.037 7.038T2 12q0 2.925 2.038 4.963T9 19Zm0-7Z"/>`),
		g.Group(children),
	)
}