package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlyFaceWhitSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSlyFaceWhitSmile0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" d="m34 21l-4-2m-11 0l-4 2m16 10s-2 4-7 4s-7-4-7-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSlyFaceWhitSmile0)"/>`),
		g.Group(children),
	)
}