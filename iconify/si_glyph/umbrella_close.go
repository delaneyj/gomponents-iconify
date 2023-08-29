package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.031 10.938v1.045h.938L7.907 1.968V.43a.411.411 0 0 0-.824 0v1.538L3.985 11.983h.953v-1.045h1.078v1.045h1.016v3.549c0 .228.251.412.479.412c.228 0 .459-.185.459-.412v-3.549h.969v-1.045h1.092z"/>`),
		g.Group(children),
	)
}