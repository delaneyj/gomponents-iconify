package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkingFaceWithOpenEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWinkingFaceWithOpenEyes0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="m33 16l-4 4l4 4m-2 7s-2 4-7 4s-7-4-7-4"/><circle cx="17" cy="20" r="4" fill="#000" stroke="#000" stroke-linecap="round"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWinkingFaceWithOpenEyes0)"/>`),
		g.Group(children),
	)
}