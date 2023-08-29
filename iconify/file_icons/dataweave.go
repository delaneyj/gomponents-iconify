package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dataweave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M408.836 54.191L179.538 381.697L53.318 201.416l49.846-71.113l99.54 142.181L256 196.33l-26.659-38.057l-26.637 38.057l-99.54-142.138L0 201.416l179.538 256.393l229.298-327.506l49.846 71.113l-126.22 180.28l-23.144-33.077l-26.66 38.077l49.804 71.113L512 201.416L408.836 54.19z"/>`),
		g.Group(children),
	)
}