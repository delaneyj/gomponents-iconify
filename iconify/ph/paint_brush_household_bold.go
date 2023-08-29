package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrushHouseholdBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M233.47 22.53a36 36 0 0 0-50.91 0c-.16.15-.3.31-.45.47L130 81l-6.22-6.21a28 28 0 0 0-39.6 0L7.52 151.51a12 12 0 0 0 0 17l80 80a12 12 0 0 0 17 0l76.7-76.7a28 28 0 0 0 0-39.6L175 126l58-52.08c.16-.15.32-.29.47-.45a36 36 0 0 0 0-50.94ZM96 223l-11-11l19.51-19.52a12 12 0 0 0-17-17L68 195l-7-7l19.51-19.52a12 12 0 0 0-17-17L44 171l-11-11l39-39l63 63ZM216.69 56.28l-60.9 54.65a20 20 0 0 0-.78 29l9.2 9.2a4 4 0 0 1 0 5.66L152 167l-63-63l12.2-12.2a4 4 0 0 1 5.66 0L116 101a20 20 0 0 0 29-.78l54.65-60.9a12 12 0 0 1 17 17Z"/>`),
		g.Group(children),
	)
}