package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25.609 6.391l.308-.573V.001h-5.824l-.583.588l-2.745 5.229l-.859.584H6.103v7.989h5.391l.479.583l-5.567 10.641l-.319.573V32h5.819l.583-.588l2.761-5.229l.853-.584h9.803V17.61h-5.401l-.479-.583l5.583-10.641z"/>`),
		g.Group(children),
	)
}