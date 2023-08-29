package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridNineThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 52H40a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12Zm-116 96v-40h56v40Zm56 8v40h-56v-40ZM36 108h56v40H36Zm64-8V60h56v40Zm64 8h56v40h-56Zm56-44v36h-56V60h52a4 4 0 0 1 4 4ZM40 60h52v40H36V64a4 4 0 0 1 4-4Zm-4 132v-36h56v40H40a4 4 0 0 1-4-4Zm180 4h-52v-40h56v36a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}