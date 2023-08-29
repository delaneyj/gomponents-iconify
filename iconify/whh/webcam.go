package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 768q-104 0-192.5-51.5t-140-140T0 384t51.5-192.5t140-140T384 0t192.5 51.5t140 140T768 384t-51.5 192.5t-140 140T384 768zm0-640q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75zm0 256q-53 0-90.5-28T256 288t37.5-68t90.5-28t90.5 28t37.5 68t-37.5 68t-90.5 28zM240 821q68 11 144 11t144-11q50 38 81 92t31 111H128q0-57 31-111.5t81-91.5z"/>`),
		g.Group(children),
	)
}