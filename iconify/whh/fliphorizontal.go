package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fliphorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 1024V0h64l384 960v64H576zM0 960L384 0h64v1024H0v-64zm384 0V160L64 960h320z"/>`),
		g.Group(children),
	)
}