package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPaperPlaneFill0"><g id="evaPaperPlaneFill1"><path id="evaPaperPlaneFill2" fill="currentColor" d="M21 4a1.31 1.31 0 0 0-.06-.27v-.09a1 1 0 0 0-.2-.3a1 1 0 0 0-.29-.19h-.09a.86.86 0 0 0-.31-.15H20a1 1 0 0 0-.3 0l-18 6a1 1 0 0 0 0 1.9l8.53 2.84l2.84 8.53a1 1 0 0 0 1.9 0l6-18A1 1 0 0 0 21 4Zm-4.7 2.29l-5.57 5.57L5.16 10ZM14 18.84l-1.86-5.57l5.57-5.57Z"/></g></g>`),
		g.Group(children),
	)
}