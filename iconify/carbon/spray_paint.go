package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SprayPaint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.5 23a4.5 4.5 0 1 1 4.5-4.5a4.505 4.505 0 0 1-4.5 4.5Zm0-7a2.5 2.5 0 1 0 2.5 2.5a2.503 2.503 0 0 0-2.5-2.5Z"/><path fill="currentColor" d="M28 8h-2V3h-7v5h-2a2.002 2.002 0 0 0-2 2v18a2.002 2.002 0 0 0 2 2h11a2.003 2.003 0 0 0 2-2V10a2.002 2.002 0 0 0-2-2zm-7-3h3v3h-3zm-4 23V10h11l.002 18zM2 14h3v3H2zm5-5h3v3H7zM2 9h3v3H2zm10-5h3v3h-3zM7 4h3v3H7zM2 4h3v3H2z"/>`),
		g.Group(children),
	)
}