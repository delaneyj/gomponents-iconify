package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 5h12v8H3zm.918 9.938H1v-2.876h1v1.98h1.918v.896zm13.082 0h-2.938v-.896H16v-1.984h1v2.88zm0-9.021h-1v-1.95h-1.943v-.946H17v2.896zM2 5.938H1V3h2.938v.938H2v2z"/>`),
		g.Group(children),
	)
}