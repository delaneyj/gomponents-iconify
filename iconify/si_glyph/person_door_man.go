package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonDoorMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.013 8.966c1.648 0 2.983-2.459 2.983-4.268c0-.239-.024-.471-.07-.695H6.098a3.53 3.53 0 0 0-.07.695c-.001 1.809 1.336 4.268 2.985 4.268zM6 1h5.943v1.927H6zm6.042 9.634v5.328h4.93s.316-5.915-4.206-5.915c-.169.233-.434.42-.724.587zM5.24 10C1.124 10 1 16 1 16h10v-4.985a10.93 10.93 0 0 1-1.986.434S6.1 11.16 5.24 10zm4.776 5.031H8.969v-1.094h1.047v1.094zm.015-2H8.953v-1.062h1.078v1.062z"/>`),
		g.Group(children),
	)
}