package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.918 13.338l-.943 1.62l-.942-1.62l.488-3.296h.908l.489 3.296zM6.534 15C4.833 12.746 5.378 9.224 5.21 9C1.123 8.999 1 15 1 15h5.534zm5.048 0H17c-.001 0 0-6.031-3.68-6.031c-.164.22-.035 3.759-1.738 6.031zM9.008 8.941C7.39 8.941 6 5.732 6 4.064c0-1.67 1.389-3.006 3.008-3.006a3.016 3.016 0 0 1 2.994 3.006c0 1.668-1.374 4.877-2.994 4.877z"/>`),
		g.Group(children),
	)
}