package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaLayer0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaLayer1" d="M 16,80 127.94695,15.974538"/><path id="galaLayer2" d="m 16,80 112,64"/><path id="galaLayer3" d="m 16,176 112,64"/><path id="galaLayer4" d="M 128,16 240,80"/><path id="galaLayer5" d="M 128,144 240,80"/><path id="galaLayer6" d="M 128,240 240,176"/><path id="galaLayer7" d="m 16,128 112,64"/><path id="galaLayer8" d="M 128,192 239.94695,128.0019"/><path id="galaLayer9" d="M 16,128 58.031909,104.01298"/><path id="galaLayera" d="M 16,176 58.032661,151.99127"/><path id="galaLayerb" d="m 239.94694,128.00191 -41.9796,-23.98708"/><path id="galaLayerc" d="M 240,176 197.96608,151.99056"/></g>`),
		g.Group(children),
	)
}