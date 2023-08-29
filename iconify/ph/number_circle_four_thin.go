package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm36-76a4 4 0 0 1-4 4h-12v28a4 4 0 0 1-8 0v-28H96a4 4 0 0 1-3.79-5.26l24-72a4 4 0 1 1 7.58 2.52L101.55 140H140v-28a4 4 0 0 1 8 0v28h12a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}