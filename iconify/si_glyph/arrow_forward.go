package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.584 5.988l-5.24-4.645c-.404-.41-.879-.479-1.283-.068v3.132l-.227-.001c-4.95 0-8.75 3.563-8.75 8.41c0 1.688.766 1.073 1.083.484c1.501-2.78 4.267-4.677 7.705-4.677l.188.001v3.065c.404.41.929.361 1.283.068l5.24-4.283c.405-.41.405-1.075.001-1.486z"/>`),
		g.Group(children),
	)
}