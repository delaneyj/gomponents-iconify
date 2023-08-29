package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DkFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#c8102e" d="M0 0h640.1v480H0z"/><path fill="#fff" d="M205.7 0h68.6v480h-68.6z"/><path fill="#fff" d="M0 205.7h640.1v68.6H0z"/>`),
		g.Group(children),
	)
}