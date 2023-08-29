package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SsFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#078930" d="M0 336h640v144H0z"/><path fill="#fff" d="M0 144h640v192H0z"/><path d="M0 0h640v144H0z"/><path fill="#da121a" d="M0 168h640v144H0z"/><path fill="#0f47af" d="m0 0l415.7 240L0 480z"/><path fill="#fcdd09" d="M200.7 194.8L61.7 240l139 45.1L114.9 167v146z"/>`),
		g.Group(children),
	)
}