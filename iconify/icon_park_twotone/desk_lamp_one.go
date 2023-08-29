package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskLampOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDeskLampOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 40.999h36m-3-22v22m-12-27l12 10"/><path fill="#555" d="M26.34 14.577a4.43 4.43 0 0 0 .567-.794c1.182-2.115.45-4.982-1.6-6.204c-2.048-1.22-4.819-.44-6.003 1.673c-.187.334-.29.595-.384.904c-3.576-.276-7.007 1.475-8.92 4.895L26.694 25c1.913-3.42 1.666-7.364-.354-10.423Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDeskLampOne0)"/>`),
		g.Group(children),
	)
}