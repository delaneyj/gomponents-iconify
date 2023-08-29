package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M243.15 116.29L139.71 12.85a16.56 16.56 0 0 0-23.42 0L12.85 116.29a16.56 16.56 0 0 0 0 23.42l103.44 103.44a16.56 16.56 0 0 0 23.42 0l103.44-103.44a16.56 16.56 0 0 0 0-23.42Zm-11.31 12.1L128.4 231.84a.58.58 0 0 1-.8 0L24.16 128.39a.56.56 0 0 1 0-.78L127.6 24.16a.58.58 0 0 1 .8 0l103.44 103.45a.56.56 0 0 1 0 .78Zm-58.18-14a8 8 0 0 1 0 11.32l-24 24a8 8 0 0 1-11.32-11.32L148.69 128H112a16 16 0 0 0-16 16v8a8 8 0 0 1-16 0v-8a32 32 0 0 1 32-32h36.69l-10.35-10.34a8 8 0 0 1 11.32-11.32Z"/>`),
		g.Group(children),
	)
}