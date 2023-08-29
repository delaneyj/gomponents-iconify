package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 6q18 0 30.5 12T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5zm0 383V91H43v298h213z"/>`),
		g.Group(children),
	)
}