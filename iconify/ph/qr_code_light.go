package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M104 42H56a14 14 0 0 0-14 14v48a14 14 0 0 0 14 14h48a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14Zm2 62a2 2 0 0 1-2 2H56a2 2 0 0 1-2-2V56a2 2 0 0 1 2-2h48a2 2 0 0 1 2 2Zm-2 34H56a14 14 0 0 0-14 14v48a14 14 0 0 0 14 14h48a14 14 0 0 0 14-14v-48a14 14 0 0 0-14-14Zm2 62a2 2 0 0 1-2 2H56a2 2 0 0 1-2-2v-48a2 2 0 0 1 2-2h48a2 2 0 0 1 2 2Zm94-158h-48a14 14 0 0 0-14 14v48a14 14 0 0 0 14 14h48a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14Zm2 62a2 2 0 0 1-2 2h-48a2 2 0 0 1-2-2V56a2 2 0 0 1 2-2h48a2 2 0 0 1 2 2Zm-64 72v-32a6 6 0 0 1 12 0v32a6 6 0 0 1-12 0Zm76-16a6 6 0 0 1-6 6h-26v42a6 6 0 0 1-6 6h-32a6 6 0 0 1 0-12h26v-58a6 6 0 0 1 12 0v10h26a6 6 0 0 1 6 6Zm0 32v16a6 6 0 0 1-12 0v-16a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}