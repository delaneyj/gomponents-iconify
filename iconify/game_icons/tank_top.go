package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TankTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m128 37l32-6c16 112 32 150 96 150s80-38 96-150l32 6s-16 160 32 208c0 96 0 112 16 224c-112 16-240 16-352 0c16-112 16-128 16-224c48-48 32-208 32-208z"/>`),
		g.Group(children),
	)
}