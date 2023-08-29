package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyXEyesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm60.24-117.76L176.49 112l11.75 11.76a6 6 0 1 1-8.48 8.48L168 120.49l-11.76 11.75a6 6 0 0 1-8.48-8.48L159.51 112l-11.75-11.76a6 6 0 0 1 8.48-8.48L168 103.51l11.76-11.75a6 6 0 0 1 8.48 8.48Zm-80 0L96.49 112l11.75 11.76a6 6 0 1 1-8.48 8.48L88 120.49l-11.76 11.75a6 6 0 0 1-8.48-8.48L79.51 112l-11.75-11.76a6 6 0 0 1 8.48-8.48L88 103.51l11.76-11.75a6 6 0 0 1 8.48 8.48ZM138 180a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}