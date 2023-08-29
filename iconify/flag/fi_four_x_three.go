package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h640v480H0z"/><path fill="#002f6c" d="M0 174.5h640v131H0z"/><path fill="#002f6c" d="M175.5 0h130.9v480h-131z"/>`),
		g.Group(children),
	)
}