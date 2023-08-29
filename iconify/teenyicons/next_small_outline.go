package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m4.5 5.5l.248-.434A.5.5 0 0 0 4 5.5h.5Zm0 4H4a.5.5 0 0 0 .748.434L4.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L8 7.5Zm-4-2v4h1v-4H4Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868ZM10 5v5h1V5h-1Z"/>`),
		g.Group(children),
	)
}