package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreatPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 28.5a112 112 0 0 0-112 112a112 112 0 0 0 24.4 69.8L256 58.49l87.5 151.71a112 112 0 0 0 24.5-69.7a112 112 0 0 0-112-112zm0 66.01l-39.8 68.99h79.6zm-50.2 86.99l-26.5 46H247v-46zm59.2 0v46h67.7l-26.5-46zm-96.1 64l-26.5 46h227.2l-26.5-46zm-36.9 64l-26.5 46H247v-46zm133 0v46h141.5l-26.5-46zm-169.93 64l-26.55 46H443.5l-26.6-46zm-36.93 64l-26.55 46H247v-46zm206.86 0v46h215.4l-26.6-46z"/>`),
		g.Group(children),
	)
}