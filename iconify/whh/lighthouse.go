package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lighthouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m768 192l256-128v320L768 256v-64zm-64 128q0 23-15 41t-38 22l12 65H361l12-65q-23-4-38-22t-15-41V128h-64L512 0l256 128h-64v192zM512 160q0-13-9.5-22.5T480 128h-64q-13 0-22.5 9.5T384 160v128q0 13 9.5 22.5T416 320h64q13 0 22.5-9.5T512 288V160zm128 0q0-13-9.5-22.5T608 128t-22.5 9.5T576 160v128q0 13 9.5 22.5T608 320t22.5-9.5T640 288V160zm-384 96L0 384V64l256 128v64zm419 256l11 64H337l12-64h326zm35 192H314l12-64h372zM287 852l16-84h418l16 84q129 24 208 70t79 102H0q0-56 79-102t208-70z"/>`),
		g.Group(children),
	)
}