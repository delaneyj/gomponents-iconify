package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 181q0 19-19 30l19 45h-32l-19-43h-24v43h-32V128h75q13 0 22.5 9.5T384 160v21zm-32 0v-21h-43v21h43zM75 171v-43h32v128H75v-53H32v53H0V128h32v43h43zm138-43q13 0 22.5 9.5T245 160v64q0 13-9.5 22.5T213 256h-74V128h74zm0 96v-64h-42v64h42z"/>`),
		g.Group(children),
	)
}