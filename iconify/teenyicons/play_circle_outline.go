package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.5 5.5l.248-.434A.5.5 0 0 0 6 5.5h.5Zm0 4H6a.5.5 0 0 0 .748.434L6.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L10 7.5ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM6 5.5v4h1v-4H6Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868Z"/>`),
		g.Group(children),
	)
}