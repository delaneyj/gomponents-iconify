package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropPortrait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 0q18 0 30.5 12.5T299 43v298q0 18-12.5 30.5T256 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h213zm0 341V43H43v298h213z"/>`),
		g.Group(children),
	)
}