package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.48 28.49a1 1 0 0 0-1.38-.32A12 12 0 0 1 12.45 22h11.71a1 1 0 0 0 0-2H11.93a11.16 11.16 0 0 1 0-4h12.23a1 1 0 0 0 0-2H12.45a12 12 0 0 1 17.61-6.2a1 1 0 0 0 1.06-1.7A14 14 0 0 0 10.34 14h-6.8a1 1 0 1 0 0 2h6.37a14 14 0 0 0-.16 2a14 14 0 0 0 .16 2H3.54a1 1 0 1 0 0 2h6.8a14 14 0 0 0 20.83 7.87a1 1 0 0 0 .31-1.38Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}