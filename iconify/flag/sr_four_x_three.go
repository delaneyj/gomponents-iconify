package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#377e3f" d="M.1 0h640v480H.1z"/><path fill="#fff" d="M.1 96h640v288H.1z"/><path fill="#b40a2d" d="M.1 144h640v192H.1z"/><path fill="#ecc81d" d="m320 153.2l56.4 173.6l-147.7-107.3h182.6L263.6 326.8z"/>`),
		g.Group(children),
	)
}