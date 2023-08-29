package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flipvertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M64 1024H0V576h1024v64zM0 0h64l960 384v64H0V0zm64 384h800L64 64v320z"/>`),
		g.Group(children),
	)
}