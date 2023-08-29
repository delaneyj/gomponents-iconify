package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclamationsign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-192q27 0 45.5-18.5t18.5-45t-19-45.5t-45-19t-45 19t-19 45.5t18.5 45T512 832zm0-640q-53 0-90.5 37.5T384 320l64 256q0 27 19 45.5t45 18.5t45-18.5t19-45.5l64-256q0-53-37.5-90.5T512 192z"/>`),
		g.Group(children),
	)
}