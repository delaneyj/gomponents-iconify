package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MqOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#231f1e" d="M0 0h512v512H0z"/><path fill="#00a650" d="M0 0h512v256H0z"/><path fill="#ef1923" d="M256 256L0 512V0z"/>`),
		g.Group(children),
	)
}