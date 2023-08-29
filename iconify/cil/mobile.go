package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M380 16H132a32.036 32.036 0 0 0-32 32v416a32.036 32.036 0 0 0 32 32h248a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32Zm0 32v32H132V48Zm0 64l.011 224H132V112Zm0 352H132v-96h248.016v96Z"/><path fill="currentColor" d="M240 400h32v32h-32z"/>`),
		g.Group(children),
	)
}