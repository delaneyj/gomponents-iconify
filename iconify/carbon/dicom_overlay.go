package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DicomOverlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M28 6v20H4V6h24m0-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z" fill="currentColor"/><path d="M6 8h10v2H6z" fill="currentColor"/><path d="M6 12h10v2H6z" fill="currentColor"/><path d="M6 16h6v2H6z" fill="currentColor"/>`),
		g.Group(children),
	)
}