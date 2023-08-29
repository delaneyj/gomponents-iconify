package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundDetectionDogBarking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.85 22v-6.775l4.925 4.925q1.35-1.35 2.413-2.912t1.062-3.463q0-1.4-.55-2.688t-1.525-2.262L8.35 6h3.5l4-4l2.05 2.05q1 1 1.55 2.263T20 9q0 1.425-.55 2.688T17.9 13.95L15.85 16v6h-11Zm4.925-4.675l-5.2-5.2q-.275-.275-.425-.65T4 10.7q0-.4.15-.763t.425-.637l2.1-2.125l3.1 3.075q.7.7 1.088 1.613t.387 1.912q0 1-.375 1.913t-1.1 1.637Z"/>`),
		g.Group(children),
	)
}