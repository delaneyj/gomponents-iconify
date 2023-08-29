package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm46.83-127.29a6 6 0 0 0-6.16.3L122 122.12V96a6 6 0 0 0-9.33-5l-48 32a6 6 0 0 0 0 10l48 32a6 6 0 0 0 9.33-5v-26.12L168.67 165a6 6 0 0 0 9.33-5V96a6 6 0 0 0-3.17-5.29ZM110 148.79L78.82 128L110 107.21Zm56 0L134.82 128L166 107.21Z"/>`),
		g.Group(children),
	)
}