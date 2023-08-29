package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperclipAttechmentTilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.536 11.465L11.43 4.57a5.25 5.25 0 1 1 7.424 7.425L10.9 19.95A3.5 3.5 0 0 1 5.95 15l7.956-7.955A1.75 1.75 0 0 1 16.38 9.52l-6.895 6.894"/>`),
		g.Group(children),
	)
}