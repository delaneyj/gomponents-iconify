package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoundationdbIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#0073E0" d="M0 123.59v-6.764h53.686V89.337l53.341 5.346v22.143h54.103V89.337L256 99.264v24.326z"/><path fill="#289DFC" d="M0 80.169L53.686 72.3V44.812l53.341 13.219v22.144L161.13 72.3V44.812L256 68.721v24.434l-94.87-13.219l-54.103 7.111l-53.341-7.111L0 87.042z"/><path fill="#9ACEFE" d="m107.027 43.905l54.103-16.417V0L256 38.942v23.671l-94.87-26.725l-54.103 14.889z"/>`),
		g.Group(children),
	)
}