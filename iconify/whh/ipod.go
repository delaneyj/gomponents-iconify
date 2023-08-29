package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024H128q-53 0-90.5-37.5T0 896V128q0-53 37.5-90.5T128 0h384q53 0 90.5 37.5T640 128v768q0 53-37.5 90.5T512 1024zm64-896q0-26-18.5-45T512 64H128q-27 0-45.5 19T64 128v192q0 27 18.5 45.5T128 384h384q26 0 45-18.5t19-45.5V128zM320 448q-106 0-181 75T64 704t75 181t181 75t181-75t75-181t-75-181t-181-75zm0 448q-80 0-136-56t-56-136t56-136t136-56t136 56t56 136t-56 136t-136 56zm0-256q-27 0-45.5 19T256 704.5t18.5 45T320 768t45.5-18.5T384 704t-18.5-45.5T320 640z"/>`),
		g.Group(children),
	)
}