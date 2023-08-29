package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchLayerThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 4v2h2.586l-5 5L21 12.414l5-5V10h2V4h-6zM10 4v2H7.414l5 5L11 12.414l-5-5V10H4V4h6zm10 1l-4-4l-4 4l1.414 1.414L15 4.829V11h2V4.829l1.586 1.585L20 5zm2 23v-2h2.586l-5-5L21 19.586l5 5V22h2v6h-6zm-12 0v-2H7.414l5-5L11 19.586l-5 5V22H4v6h6zm10-1l-4 4l-4-4l1.414-1.414L15 27.171V21h2v6.171l1.586-1.585L20 27zM5 12l-4 4l4 4l1.414-1.414L4.829 17H11v-2H4.829l1.585-1.586L5 12zm22 0l4 4l-4 4l-1.414-1.414L27.171 17H21v-2h6.171l-1.585-1.586L27 12z"/>`),
		g.Group(children),
	)
}