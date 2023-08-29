package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBotswana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M5.133 39.867a27.992 27.992 0 0 1-.692-2.95h55.118a27.742 27.742 0 0 1-.691 2.95H5.133m-.692-12.784c.179-1.003.409-1.986.692-2.949h53.734c.283.963.514 1.946.691 2.949H4.441"/>`),
		g.Group(children),
	)
}