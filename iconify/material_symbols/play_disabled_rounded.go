package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.55 18.025q-.5.325-1.025.038T8 17.175V10.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.688t-.3.712q-.3.275-.7.275t-.7-.3L13 15.8l-3.45 2.225Zm6.9-4.425L8 5.15V5l9.675 6.15q.475.275.463.838t-.488.862l-1.2.75Z"/>`),
		g.Group(children),
	)
}