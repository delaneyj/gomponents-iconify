package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Popup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1091 0v636H818v364H0V364h273V0h818zM818 545h182V182H364v182h454v181zM91 909h636V545H91v364z"/>`),
		g.Group(children),
	)
}