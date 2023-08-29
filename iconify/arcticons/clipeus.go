package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipeus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="34.383" height="22.041" x="6.808" y="12.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.5" ry="3.5" transform="rotate(-45 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.793 31.793L16.207 16.207"/>`),
		g.Group(children),
	)
}