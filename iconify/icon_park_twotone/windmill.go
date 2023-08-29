package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windmill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWindmill0"><path fill="#555" fill-rule="evenodd" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 23.992l-.03-9.996L23.943 4L12 14v10l12-.008Zm.008.008l9.996-.03L44 23.943L34 12H24l.008 12Zm-.008.008l.03 9.996l.028 9.996L36 34V24l-12 .008ZM23.992 24l-9.996.03L4 24.057L14 36h10l-.008-12Z" clip-rule="evenodd"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWindmill0)"/>`),
		g.Group(children),
	)
}