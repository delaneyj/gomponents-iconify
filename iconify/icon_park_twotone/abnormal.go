package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abnormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAbnormal0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 16.398V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h10"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M16 14h13m-13 7h5"/><circle cx="34" cy="34" r="10" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" transform="rotate(90 34 34)"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M34 36v3"/><circle cx="34" cy="30" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAbnormal0)"/>`),
		g.Group(children),
	)
}