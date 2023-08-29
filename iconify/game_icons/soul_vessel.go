package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoulVessel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 32A224 224 0 0 0 32 256a224 224 0 0 0 224 224a224 224 0 0 0 224-224A224 224 0 0 0 256 32zM132.1 282.8c25.2.4 47 17.3 58.2 27.6c49.3 45.3 16.4 87.4-2.7 96.6c-33.9 16.5-68.9 9.7-86.8-16.1c-36.65-52.9-17.55-89.4 1.6-100.3c9.3-5.3 18.4-7.5 27.2-7.8h2.5zm247.8 0h2.5c8.8.3 17.9 2.5 27.2 7.8c19.2 10.9 38.3 47.4 1.6 100.3c-17.9 25.8-52.9 32.6-86.8 16.1c-19.1-9.2-52-51.3-2.7-96.6c11.2-10.3 33-27.2 58.2-27.6z"/>`),
		g.Group(children),
	)
}