package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 896v64q0 27-19 45.5t-45 18.5H192q-27 0-45.5-18.5T128 960v-64q-53 0-90.5-37.5T0 768V512q0-53 37.5-90.5T128 384h768q53 0 90.5 37.5T1024 512v256q0 53-37.5 90.5T896 896zm0-320q0-26-19-45t-45-19H192q-27 0-45.5 19T128 576v128q0 27 18.5 45.5T192 768h640q26 0 45-18.5t19-45.5V576zM256 704q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45-19t45.5 19t19 45.5t-18.5 45T256 704zM128 64q0-26 18.5-45T192 0h640q26 0 45 19t19 45v256H128V64z"/>`),
		g.Group(children),
	)
}