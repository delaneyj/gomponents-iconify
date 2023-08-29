package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Erroralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm.5 512q-26.5 0-45.5-19t-19-45l-64-256q0-53 37.5-90.5T512 192t90.5 37.5T640 320l-64 256q0 26-18.5 45t-45 19zm-.5 64q27 0 45.5 18.5t18.5 45t-19 45.5t-45.5 19t-45-19t-18.5-45.5t18.5-45T512 704z"/>`),
		g.Group(children),
	)
}