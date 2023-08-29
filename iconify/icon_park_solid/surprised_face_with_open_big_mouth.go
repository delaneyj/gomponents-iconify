package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurprisedFaceWithOpenBigMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSurprisedFaceWithOpenBigMouth0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="M31 18v1m-14-1v1"/><rect width="12" height="12" x="18" y="24" fill="#000" stroke="#000" stroke-linecap="round" rx="6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSurprisedFaceWithOpenBigMouth0)"/>`),
		g.Group(children),
	)
}