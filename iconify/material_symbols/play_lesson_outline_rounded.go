package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayLessonOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.3 20q.15.5.413 1.038t.537.962H5q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h12q.825 0 1.413.588T19 4v7.1q-.45-.05-1-.05t-1 .05V4h-5v6.125q0 .3-.25.438t-.5-.013L9.5 9.5l-1.75 1.05q-.25.15-.5.013T7 10.124V4H5v16h6.3Zm6.7 3q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm-.475-2.975l2.55-1.6q.25-.15.25-.425t-.25-.425l-2.55-1.6q-.25-.15-.513-.013t-.262.438v3.2q0 .3.263.438t.512-.013ZM7 4h5h-5Zm4.3 0H5h12h-6h.3Z"/>`),
		g.Group(children),
	)
}