package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.016.06a7.97 7.97 0 1 0 .002 15.936A7.97 7.97 0 0 0 9.016.06zM3.049 8.028a5.974 5.974 0 0 1 5.967-5.966c1.354 0 2.6.458 3.602 1.221l-8.347 8.348a5.931 5.931 0 0 1-1.222-3.603zm5.967 5.966a5.921 5.921 0 0 1-3.447-1.105l8.309-8.309a5.93 5.93 0 0 1 1.104 3.448a5.974 5.974 0 0 1-5.966 5.966z"/>`),
		g.Group(children),
	)
}