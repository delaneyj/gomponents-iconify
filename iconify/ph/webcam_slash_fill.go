package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.92 210.62L158 149.12L53.92 34.62a8 8 0 1 0-11.84 10.76L58.82 63.8A80 80 0 0 0 120 183.6V200H32a8 8 0 0 0 0 16h165.19l4.89 5.38a8 8 0 1 0 11.84-10.76ZM80 104a48.19 48.19 0 0 1 2.24-14.44L96 104.71a32 32 0 0 0 28.23 31.06L138 150.94A48 48 0 0 1 80 104Zm56 96v-16.41a79.91 79.91 0 0 0 25.44-6.91l21.2 23.32ZM86.56 46.74a8 8 0 0 1 2-12.34a80 80 0 0 1 105 115.5a8 8 0 0 1-6.05 3.39h-.5a8 8 0 0 1-5.92-2.62L170 138.46a8 8 0 0 1-.92-9.54a48 48 0 0 0-61.95-68.14A8 8 0 0 1 97.67 59ZM128 72a32 32 0 0 1 32 32a35.18 35.18 0 0 1-.19 3.55a8 8 0 0 1-13.88 4.5l-24.24-26.67a8 8 0 0 1 5.8-13.38Z"/>`),
		g.Group(children),
	)
}