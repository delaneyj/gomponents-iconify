package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.937 8.034c-.283 0-.552.055-.798.154C12.975 6.401 11.416 5 9.514 5c-.465 0-.917.088-1.317.237c-.156.058-.197.117-.197.233v6.292c0 .121.098.222.221.234l5.717.003c1.139 0 2.062-.888 2.062-1.983s-.924-1.983-2.063-1.983zM6.25 12h.5L7 8.497L6.75 5h-.5L6 8.497zm-1.5 0h-.5L4 9.457L4.25 7h.5L5 9.5zm-2.5 0h.5L3 10l-.25-2h-.5L2 10zm-2-1h.5L1 10L.75 9h-.5L0 10z"/>`),
		g.Group(children),
	)
}