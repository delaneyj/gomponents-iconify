package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-630-359l213 217q14 14 33 14t33-14l213-217q12-13 8-19t-19-6h-107V512q0-27-18.5-45.5t-45.5-18.5h-128q-27 0-45.5 18.5t-18.5 45.5v128h-108q-15 0-18.5 6t8.5 19zm630-505q0-13-9.5-22.5t-22.5-9.5h-704q-13 0-22.5 9.5t-9.5 22.5v128q0 13 9.5 22.5t22.5 9.5h704q13 0 22.5-9.5t9.5-22.5V160z"/>`),
		g.Group(children),
	)
}