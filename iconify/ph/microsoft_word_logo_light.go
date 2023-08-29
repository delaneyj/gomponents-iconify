package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftWordLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 26H72a14 14 0 0 0-14 14v26H40a14 14 0 0 0-14 14v96a14 14 0 0 0 14 14h18v26a14 14 0 0 0 14 14h128a14 14 0 0 0 14-14V40a14 14 0 0 0-14-14Zm-42 76h44v52h-44ZM70 40a2 2 0 0 1 2-2h128a2 2 0 0 1 2 2v50h-44V80a14 14 0 0 0-14-14H70ZM38 176V80a2 2 0 0 1 2-2h104a2 2 0 0 1 2 2v96a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2Zm162 42H72a2 2 0 0 1-2-2v-26h74a14 14 0 0 0 14-14v-10h44v50a2 2 0 0 1-2 2ZM70.18 153.46l-12-48a6 6 0 1 1 11.64-2.92l8.07 32.27l8.74-17.49a6 6 0 0 1 10.74 0l8.74 17.49l8.07-32.27a6 6 0 1 1 11.64 2.92l-12 48a6 6 0 0 1-5.17 4.5a4.63 4.63 0 0 1-.65 0a6 6 0 0 1-5.37-3.32L92 133.42l-10.63 21.26a6 6 0 0 1-11.19-1.22Z"/>`),
		g.Group(children),
	)
}