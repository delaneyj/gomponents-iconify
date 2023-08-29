package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#c1272d" d="M512 0H0v512h512z"/><path fill="none" stroke="#006233" stroke-width="12.5" d="m256 191.4l-38 116.8l99.4-72.2H194.6l99.3 72.2z"/>`),
		g.Group(children),
	)
}