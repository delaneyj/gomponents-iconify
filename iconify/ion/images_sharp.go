package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<ellipse cx="373.14" cy="219.33" fill="none" rx="46.29" ry="46"/><path fill="currentColor" d="M80 132v328a20 20 0 0 0 20 20h392a20 20 0 0 0 20-20V132a20 20 0 0 0-20-20H100a20 20 0 0 0-20 20Zm293.14 41.33a46 46 0 1 1-46.28 46a46.19 46.19 0 0 1 46.28-46Zm-261.41 276v-95.48l122.76-110.2L328.27 337l-113 112.33Zm368.27 0H259l144.58-144L480 370.59Z"/><path fill="currentColor" d="M20 32A20 20 0 0 0 0 52v344a20 20 0 0 0 20 20h28V100a20 20 0 0 1 20-20h380V52a20 20 0 0 0-20-20Z"/>`),
		g.Group(children),
	)
}