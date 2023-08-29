package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#D91023" d="M0 0h512v512H0z"/><path fill="#fff" d="M170.7 0h170.6v512H170.7z"/>`),
		g.Group(children),
	)
}