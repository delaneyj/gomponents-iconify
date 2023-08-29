package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MvFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#d21034" d="M0 0h640v480H0z"/><path fill="#007e3a" d="M120 120h400v240H120z"/><circle cx="350" cy="240" r="80" fill="#fff"/><circle cx="380" cy="240" r="80" fill="#007e3a"/>`),
		g.Group(children),
	)
}