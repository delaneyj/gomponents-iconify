package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForHonduras(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#42ade2" d="M32 62c13.1 0 24.2-8.3 28.3-20H3.7C7.8 53.7 18.9 62 32 62zm0-60C18.9 2 7.8 10.3 3.7 22h56.6C56.2 10.3 45.1 2 32 2z"/><path fill="#f9f9f9" d="M60.3 42c1.1-3.1 1.7-6.5 1.7-10s-.6-6.9-1.7-10H3.7C2.6 25.1 2 28.5 2 32s.6 6.9 1.7 10h56.6z"/><path fill="#42ade2" d="m32 33.6l1.9 1.4l-.7-2.3l1.8-1.4h-2.3L32 29l-.7 2.3H29l1.8 1.4l-.7 2.3zm-10-5l1.9 1.4l-.7-2.3l1.8-1.4h-2.3L22 24l-.7 2.3H19l1.8 1.4l-.7 2.3zm20 0l1.9 1.4l-.7-2.3l1.8-1.4h-2.3L42 24l-.7 2.3H39l1.8 1.4l-.7 2.3zm-20 10l1.9 1.4l-.7-2.3l1.8-1.4h-2.3L22 34l-.7 2.3H19l1.8 1.4l-.7 2.3zm20 0l1.9 1.4l-.7-2.3l1.8-1.4h-2.3L42 34l-.7 2.3H39l1.8 1.4l-.7 2.3z"/>`),
		g.Group(children),
	)
}