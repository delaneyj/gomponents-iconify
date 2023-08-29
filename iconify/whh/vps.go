package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 384h-320v128h384q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5h-896q-26 0-45-18.5t-19-45t19-45.5t45-19h384V384h-320q-53 0-90.5-37.5T.428 256V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v128q0 53-37.5 90.5t-90.5 37.5zm-704-256q-27 0-45.5 19t-18.5 45.5t18.5 45t45 18.5t45.5-18.5t19-45t-18.5-45.5t-45.5-19zm640 0h-384q-26 0-45 19t-19 45.5t19 45t45 18.5h384q27 0 45.5-18.5t18.5-45t-18.5-45.5t-45.5-19zm-320 640q104 0 192.5 17t140 46.5t51.5 64.5t-51.5 64.5t-140 46.5t-192.5 17t-192.5-17t-140-46.5t-51.5-64.5t51.5-64.5t140-46.5t192.5-17z"/>`),
		g.Group(children),
	)
}