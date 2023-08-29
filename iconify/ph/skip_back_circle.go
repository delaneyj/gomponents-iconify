package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm35.88-135a8 8 0 0 0-8.12.22L104 113.57V88a8 8 0 0 0-16 0v80a8 8 0 0 0 16 0v-25.57l51.76 32.35A8 8 0 0 0 168 168V88a8 8 0 0 0-4.12-7ZM152 153.57L111.09 128L152 102.43Z"/>`),
		g.Group(children),
	)
}