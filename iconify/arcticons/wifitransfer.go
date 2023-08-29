package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifitransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.84 13.818a7.575 7.575 0 0 0-9.674-.024v.024m12.182-3.05a11.535 11.535 0 0 0-14.695 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.675 7.995a15.173 15.173 0 0 0-19.393.015M36.51 43.5H11.49a.715.715 0 0 1-.714-.715V17.767a.715.715 0 0 1 .715-.715h22.637l3.096 3.096v22.637a.715.715 0 0 1-.715.715Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.993 29.204h20.015V43.5H13.993zm4.729-12.152h13.581v7.863H18.722z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.474 18.809h3.183v4.349h-3.183zM16.277 32.838h15.446m0 3.514H16.277m0 3.514h15.446"/>`),
		g.Group(children),
	)
}