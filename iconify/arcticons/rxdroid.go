package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rxdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="43.995" height="20.112" x="2.002" y="13.944" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.056" transform="rotate(-45 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.111 31.111L16.889 16.889"/>`),
		g.Group(children),
	)
}