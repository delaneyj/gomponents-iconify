package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16 17.06l4.72 4.72a.75.75 0 1 0 1.06-1.06L3.28 2.22a.75.75 0 0 0-1.06 1.06l4.722 4.723l-3.456 3.46a3.25 3.25 0 0 0 .004 4.596l4.462 4.455a3.255 3.255 0 0 0 4.596-.001L16 17.06Zm5.05-5.05L18.06 15L9.002 5.94l2.984-2.988a3.25 3.25 0 0 1 2.3-.953h5.465A2.25 2.25 0 0 1 22 4.25v5.462a3.25 3.25 0 0 1-.952 2.298ZM17 5.502a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}