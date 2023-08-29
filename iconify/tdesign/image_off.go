package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 .586L.586 22L2 23.414L3.414 22H22V3.414L23.414 2L22 .586ZM11.5 13.914L17.586 20H5.414l6.086-6.086Zm8.5 5.671L12.915 12.5l1.06-1.062a3.21 3.21 0 0 0 4.463-4.463L20 5.414v14.172Zm-4.47-9.613l1.443-1.442a1.21 1.21 0 0 1-1.442 1.443ZM2.004 2l-.005 15.658l2.005-1.696V4l11.362.002l2.37-2.004L2.003 2Z"/>`),
		g.Group(children),
	)
}