package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 72a56 56 0 0 0-39.14 96H95.14A56 56 0 1 0 56 184h144a56 56 0 0 0 0-112ZM56 168a40 40 0 1 1 40-40a40 40 0 0 1-40 40Zm144 0a40 40 0 1 1 40-40a40 40 0 0 1-40 40Zm24-40a24 24 0 1 1-24-24a24 24 0 0 1 24 24Zm-144 0a24 24 0 1 1-24-24a24 24 0 0 1 24 24Z"/>`),
		g.Group(children),
	)
}