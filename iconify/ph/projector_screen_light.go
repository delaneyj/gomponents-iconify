package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 78a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14H40a14 14 0 0 0-14 14v16a14 14 0 0 0 14 14h2v100H32a6 6 0 0 0 0 12h90v20.84a22 22 0 1 0 12 0V190h90a6 6 0 0 0 0-12h-10V78Zm-78 154a10 10 0 1 1-10-10a10 10 0 0 1 10 10ZM38 64V48a2 2 0 0 1 2-2h176a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2Zm164 114H54V78h148Z"/>`),
		g.Group(children),
	)
}