package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#c1272d" d="M640 0H0v480h640z"/><path fill="none" stroke="#006233" stroke-width="11.7" d="M320 179.4L284.4 289l93.2-67.6H262.4l93.2 67.6z"/>`),
		g.Group(children),
	)
}