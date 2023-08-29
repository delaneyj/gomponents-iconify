package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandCardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33 6H3a1 1 0 0 0-1 1v22a1 1 0 0 0 1 1h30a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1Zm-1 22H4V8h28Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M13.48 15.86L18 11.34l4.52 4.52a.77.77 0 0 0 .56.24a.81.81 0 0 0 .57-1.37L18 9.08l-5.65 5.65a.8.8 0 1 0 1.13 1.13Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M13.48 21.86L18 17.34l4.52 4.52a.77.77 0 0 0 .56.24a.81.81 0 0 0 .57-1.37L18 15.08l-5.65 5.65a.8.8 0 1 0 1.13 1.13Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}