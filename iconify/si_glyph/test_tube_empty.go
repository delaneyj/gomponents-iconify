package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1V.023H6V1h1.012v6L3 15s0 .962 1 .962h10c1 0 1-.962 1-.962l-4.042-8l-.02-6H12zm2 14.031H4L8 7V1h2v6l4 8.031z"/>`),
		g.Group(children),
	)
}