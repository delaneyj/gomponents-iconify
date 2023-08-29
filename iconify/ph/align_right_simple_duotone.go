package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 96v64a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8V96a8 8 0 0 1 8-8h152a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M232 56v144a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0Zm-32 40v64a16 16 0 0 1-16 16H32a16 16 0 0 1-16-16V96a16 16 0 0 1 16-16h152a16 16 0 0 1 16 16Zm-16 0H32v64h152Z"/></g>`),
		g.Group(children),
	)
}