package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.38 221.92a8 8 0 0 1-11.3-.54l-26.46-29.1A79.74 79.74 0 0 1 136 207.59V232a8 8 0 0 1-16 0v-24.4A80.11 80.11 0 0 1 48 128a8 8 0 0 1 16 0a64 64 0 0 0 100.79 52.36l-10.88-12A48 48 0 0 1 80 128V87.09L42.08 45.38a8 8 0 1 1 11.84-10.76l160 176a8 8 0 0 1-.54 11.3Zm-51.3-92.11a8 8 0 0 0 13.92-5.38V64a48 48 0 0 0-88.84-25.22a8 8 0 0 0 .84 9.59Zm30.1 31.83a8 8 0 0 0 10.36-4.55A79.62 79.62 0 0 0 208 128a8 8 0 0 0-16 0a63.71 63.71 0 0 1-4.36 23.27a8 8 0 0 0 4.54 10.37Z"/>`),
		g.Group(children),
	)
}