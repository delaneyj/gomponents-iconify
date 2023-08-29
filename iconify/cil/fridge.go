package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M376 16H136a48.054 48.054 0 0 0-48 48v400a32.036 32.036 0 0 0 32 32h272a32.036 32.036 0 0 0 32-32V64a48.054 48.054 0 0 0-48-48Zm16 448H120V240h272Zm0-256H120V64a16.019 16.019 0 0 1 16-16h240a16.019 16.019 0 0 1 16 16Z"/><path fill="currentColor" d="M144 280h32v96h-32zm0-176h32v64h-32z"/>`),
		g.Group(children),
	)
}