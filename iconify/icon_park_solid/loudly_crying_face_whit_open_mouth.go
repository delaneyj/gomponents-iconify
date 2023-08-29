package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoudlyCryingFaceWhitOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLoudlyCryingFaceWhitOpenMouth0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="M24 29c5 0 7 4 7 4H17s2-4 7-4Zm11-11l-7-1m5 1v9M20 17l-7 1m2 0v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLoudlyCryingFaceWhitOpenMouth0)"/>`),
		g.Group(children),
	)
}