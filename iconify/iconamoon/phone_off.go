package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m14.83 9.174l-1.519-1.52a9.94 9.94 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96a15.903 15.903 0 0 1-1.721 4.168m-5.3.357a12.03 12.03 0 0 1-3.311 2.346l-1.52-1.52a9.94 9.94 0 0 1-2.656.64C4.448 14.005 4 14.448 4 15v1c0 2.21 1.819 4.051 3.96 3.506a15.98 15.98 0 0 0 7.354-4.192"/><path d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}