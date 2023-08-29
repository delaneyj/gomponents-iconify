package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseRolling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 56c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8v72h-96V56zm176 72h-32V56c0-30.9-25.1-56-56-56h-80c-30.9 0-56 25.1-56 56v72H64c-35.3 0-64 28.7-64 64v224c0 35.3 28.7 64 64 64c0 17.7 14.3 32 32 32s32-14.3 32-32h128c0 17.7 14.3 32 32 32s32-14.3 32-32c35.3 0 64-28.7 64-64V192c0-35.3-28.7-64-64-64zm-208 96h160c8.8 0 16 7.2 16 16s-7.2 16-16 16H112c-8.8 0-16-7.2-16-16s7.2-16 16-16zm0 128h160c8.8 0 16 7.2 16 16s-7.2 16-16 16H112c-8.8 0-16-7.2-16-16s7.2-16 16-16z"/>`),
		g.Group(children),
	)
}