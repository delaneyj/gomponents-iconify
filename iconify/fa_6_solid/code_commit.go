package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 336a80 80 0 1 0 0-160a80 80 0 1 0 0 160zm156.8-48C462 361 397.4 416 320 416s-142-55-156.8-128H32c-17.7 0-32-14.3-32-32s14.3-32 32-32h131.2C178 151 242.6 96 320 96s142 55 156.8 128H608c17.7 0 32 14.3 32 32s-14.3 32-32 32H476.8z"/>`),
		g.Group(children),
	)
}