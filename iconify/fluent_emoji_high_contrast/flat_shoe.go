package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlatShoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 28h1.97a1.32 1.32 0 0 1 1.28 1c.147.588.675 1 1.28 1H29.5c.276 0 .5-.224.498-.5c-.016-1.724-.19-3.096-1.998-4c-1.318-.66-2.854-1.101-3.747-1.326a1.867 1.867 0 0 0-1.416.224L21 25.5l-1.152.288a3 3 0 0 1-1.677-.064L4.142 21.047A.923.923 0 0 0 3.85 21A1.85 1.85 0 0 0 2 22.85v6.65a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V28Z"/>`),
		g.Group(children),
	)
}