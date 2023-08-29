package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .186l3.617 3.082l4.737.378l.378 4.737L23.814 12l-3.082 3.617l-.378 4.737l-4.737.378L12 23.814l-3.616-3.082l-4.737-.378l-.378-4.737L.187 12l3.082-3.617l.378-4.737l4.737-.378L12 .186Zm0 2.628L9.188 5.21l-3.683.293l-.294 3.684L2.814 12l2.397 2.812l.294 3.684l3.683.294L12 21.186l2.813-2.396l3.683-.294l.294-3.684L21.186 12L18.79 9.188l-.294-3.684l-3.683-.293L12 2.814ZM17.915 9.5L11 16.414L6.586 12L8 10.586l3 3l5.5-5.5L17.915 9.5Z"/>`),
		g.Group(children),
	)
}