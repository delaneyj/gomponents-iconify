package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m384 107l85 85h-64q0 71-50 121t-120 50q-49 0-91-27l31-31q27 15 60 15q53 0 90.5-37.5T363 192h-64zm-277 85h64l-86 85l-85-85h64q0-71 50-121t121-50q49 0 91 27l-32 31q-27-15-59-15q-53 0-90.5 37.5T107 192z"/>`),
		g.Group(children),
	)
}