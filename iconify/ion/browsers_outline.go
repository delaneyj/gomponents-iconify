package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowsersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="416" height="384" x="48" y="64" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="48" ry="48"/><path fill="currentColor" d="M397.82 64H114.18C77.69 64 48 94.15 48 131.2V176h16c0-16 16-32 32-32h320c16 0 32 16 32 32h16v-44.8c0-37.05-29.69-67.2-66.18-67.2Z"/>`),
		g.Group(children),
	)
}