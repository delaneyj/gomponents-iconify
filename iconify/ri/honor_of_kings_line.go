package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HonorOfKingsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.33 4.256l-1.424 1.423a8.001 8.001 0 0 0-12.272 9.444l2.417-2.417a5.002 5.002 0 0 1 7.707-4.879l-1.464 1.465a3.001 3.001 0 0 0-4.001 4l-6.45 6.45c-.034-3.5-.591-4.811-.788-6.701A9.98 9.98 0 0 1 4.93 4.929c3.666-3.666 9.471-3.89 13.4-.673Zm2.83.002c.033 3.5.59 4.81.787 6.701a9.98 9.98 0 0 1-2.875 8.112c-3.666 3.666-9.471 3.89-13.4.673l1.424-1.423a8.001 8.001 0 0 0 12.272-9.444l-2.417 2.417a5.002 5.002 0 0 1-7.707 4.878l1.464-1.464a3.001 3.001 0 0 0 4.001-4l6.45-6.45Z"/>`),
		g.Group(children),
	)
}