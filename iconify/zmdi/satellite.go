package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zM43 42v65q26 0 45-19t19-46H43zm0 150q62 0 105.5-44T192 42h-43q0 45-31 76t-75 31v43zm0 128h298l-96-128l-74 96l-54-64z"/>`),
		g.Group(children),
	)
}