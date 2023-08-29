package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MqFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#231f1e" d="M0 0h640v480H0z"/><path fill="#00a650" d="M0 0h640v240H0z"/><path fill="#ef1923" d="m0 0l320 240L0 480z"/>`),
		g.Group(children),
	)
}