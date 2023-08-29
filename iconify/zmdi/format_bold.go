package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M183 166q21 10 33.5 29.5T229 239q0 34-23 57.5T150 320H0V21h133q36 0 61 25t25 61q0 35-36 59zM64 75v64h64q13 0 22.5-9.5t9.5-23t-9.5-22.5t-22.5-9H64zm75 192q13 0 22.5-9.5t9.5-23t-9.5-22.5t-22.5-9H64v64h75z"/>`),
		g.Group(children),
	)
}