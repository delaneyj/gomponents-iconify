package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CpuSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32.2 23.55H30v-5.4h2.2a.8.8 0 1 0 0-1.6H30V15.4h-5v8.41A1.18 1.18 0 0 1 24 25H13v-2h10v-7.6h-.77A3.68 3.68 0 0 1 19 9.89L21.29 6h-1.94V3.8a.8.8 0 1 0-1.6 0V6h-5.4V3.8a.8.8 0 1 0-1.6 0V6H8.1A2.1 2.1 0 0 0 6 8.1v1.45H3.8a.8.8 0 1 0 0 1.6H6v5.4H3.8a.8.8 0 1 0 0 1.6H6v5.4H3.8a.8.8 0 1 0 0 1.6H6v2.75A2.1 2.1 0 0 0 8.1 30h2.65v2.2a.8.8 0 1 0 1.6 0V30h5.4v2.2a.8.8 0 1 0 1.6 0V30h5.4v2.2a.8.8 0 1 0 1.6 0V30h1.55a2.1 2.1 0 0 0 2.1-2.1v-2.75h2.2a.8.8 0 1 0 0-1.6Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-3--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}