package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightAlarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12h11v2H2zm10-1H3V9.667C3 7.826 3.75 5 7.5 5S12 7.826 12 9.667V11zM7 2h1v2H7zM2.597 3.44l.707-.708l1.414 1.414l-.707.707zm10.413-.293l.707.707l-1.414 1.414l-.707-.707zM13 8h2v1h-2zM0 8h1.75v1H0z"/>`),
		g.Group(children),
	)
}