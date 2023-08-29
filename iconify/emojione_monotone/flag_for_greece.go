package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGreece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm21.677 12.299H32v5.9h25.383a27.699 27.699 0 0 1 1.986 5.9H32v-5.9H20.2V32H60c0 2.023-.221 3.996-.631 5.9H4.631A27.997 27.997 0 0 1 4 32h10.3V20.2H6.617a28.023 28.023 0 0 1 3.707-5.9H14.3v-3.976a28.064 28.064 0 0 1 5.9-3.707V14.3H32V8.4h15.038a28.186 28.186 0 0 1 6.639 5.899zM10.323 49.7a28.104 28.104 0 0 1-3.707-5.9h50.768a28.06 28.06 0 0 1-3.706 5.9H10.323zm6.638 5.9h30.078A27.82 27.82 0 0 1 32 60a27.82 27.82 0 0 1-15.039-4.4z"/>`),
		g.Group(children),
	)
}