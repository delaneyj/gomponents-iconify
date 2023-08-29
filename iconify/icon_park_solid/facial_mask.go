package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacialMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFacialMask0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M10 15.114a3 3 0 0 1 1.991-2.825l11-3.929a3 3 0 0 1 2.018 0l11 3.929A3 3 0 0 1 38 15.114v16.49c0 .885-.39 1.725-1.092 2.264C34.468 35.742 28.434 40 24 40c-4.434 0-10.467-4.258-12.908-6.132A2.838 2.838 0 0 1 10 31.604v-16.49Z"/><path stroke="#fff" stroke-linecap="round" d="M10 28c-3.314 0-6-3.11-6-6.947S6.686 16 10 16m28 12c3.314 0 6-3.11 6-6.947S41.314 16 38 16"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m16 25l8-4l8 4m-13 6l5-2l5 2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFacialMask0)"/>`),
		g.Group(children),
	)
}