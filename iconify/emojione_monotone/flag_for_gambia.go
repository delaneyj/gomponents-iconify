package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGambia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M6.175 43a29.078 29.078 0 0 1-.734-2h53.117a29.078 29.078 0 0 1-.734 2H6.175m-.734-20c.223-.666.464-1.357.734-2h51.649c.271.643.512 1.334.734 2H5.441"/>`),
		g.Group(children),
	)
}