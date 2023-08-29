package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 74a54 54 0 0 0-33.89 96H89.89A54 54 0 1 0 56 182h144a54 54 0 0 0 0-108ZM14 128a42 42 0 1 1 42 42a42 42 0 0 1-42-42Zm186 42a42 42 0 1 1 42-42a42 42 0 0 1-42 42Z"/>`),
		g.Group(children),
	)
}