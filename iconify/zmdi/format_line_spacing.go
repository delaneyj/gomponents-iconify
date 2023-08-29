package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M96 85v214h53l-74 74l-75-74h53V85H0l75-74l74 74H96zm85-42h256v42H181V43zm0 298v-42h256v42H181zm0-128v-42h256v42H181z"/>`),
		g.Group(children),
	)
}