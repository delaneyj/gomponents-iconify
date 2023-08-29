package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infographic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M915 826L730 641q37-61 37-131q0-89-54.5-158T574 261V0q126 16 228.5 87t162 182.5T1024 510q0 176-109 316zM262 445H0q22-174 146.5-298.5T445 0v261q-67 18-116.5 67.5T262 445zm-1 129q22 84 91 138.5T510 767q70 0 131-37l185 185q-140 109-316 109q-129 0-240.5-59.5T87 802.5T0 574h261z"/>`),
		g.Group(children),
	)
}