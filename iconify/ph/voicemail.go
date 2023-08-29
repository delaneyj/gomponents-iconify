package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 72a56 56 0 0 0-39.14 96H95.14A56 56 0 1 0 56 184h144a56 56 0 0 0 0-112ZM16 128a40 40 0 1 1 40 40a40 40 0 0 1-40-40Zm184 40a40 40 0 1 1 40-40a40 40 0 0 1-40 40Z"/>`),
		g.Group(children),
	)
}