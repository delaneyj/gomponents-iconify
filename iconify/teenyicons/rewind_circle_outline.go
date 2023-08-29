package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.5 5.5H7a.5.5 0 0 0-.748-.434L6.5 5.5Zm0 4l-.248.434A.5.5 0 0 0 7 9.5h-.5ZM3 7.5l-.248-.434a.5.5 0 0 0 0 .868L3 7.5Zm7.5-2h.5a.5.5 0 0 0-.748-.434l.248.434Zm0 4l-.248.434A.5.5 0 0 0 11 9.5h-.5ZM7 7.5l-.248-.434a.5.5 0 0 0 0 .868L7 7.5Zm.5 7.5A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm0 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM6 5.5v4h1v-4H6Zm.748 3.566l-3.5-2l-.496.868l3.5 2l.496-.868Zm-3.5-1.132l3.5-2l-.496-.868l-3.5 2l.496.868ZM10 5.5v4h1v-4h-1Zm.748 3.566l-3.5-2l-.496.868l3.5 2l.496-.868Zm-3.5-1.132l3.5-2l-.496-.868l-3.5 2l.496.868Z"/>`),
		g.Group(children),
	)
}