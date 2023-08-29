package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 16h480v32H16zm0 480h480V368H16Zm32-96h416v64H48ZM416 96H96v37.86l159.923 169.364L416 135.921ZM256.077 256.776L134.478 128h244.813Z"/>`),
		g.Group(children),
	)
}