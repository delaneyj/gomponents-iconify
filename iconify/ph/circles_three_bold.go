package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 76a48 48 0 1 0-48 48a48.05 48.05 0 0 0 48-48Zm-48 24a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm60 24a48 48 0 1 0 48 48a48.05 48.05 0 0 0-48-48Zm0 72a24 24 0 1 1 24-24a24 24 0 0 1-24 24ZM68 124a48 48 0 1 0 48 48a48.05 48.05 0 0 0-48-48Zm0 72a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/>`),
		g.Group(children),
	)
}