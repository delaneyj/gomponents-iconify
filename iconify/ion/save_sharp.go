package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M380.44 32H64a32 32 0 0 0-32 32v384a32 32 0 0 0 32 32h384a32.09 32.09 0 0 0 32-32V131.56ZM112 176v-64h192v64Zm223.91 179.76a80 80 0 1 1-83.66-83.67a80.21 80.21 0 0 1 83.66 83.67Z"/>`),
		g.Group(children),
	)
}