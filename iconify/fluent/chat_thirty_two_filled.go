package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 16C2 8.268 8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14c-2.368 0-4.602-.589-6.56-1.629l-5.528 1.572A1.5 1.5 0 0 1 2.06 28.09l1.572-5.527A13.943 13.943 0 0 1 2 16Zm8-3a1 1 0 0 0 1 1h10a1 1 0 1 0 0-2H11a1 1 0 0 0-1 1Zm1 5a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6Z"/>`),
		g.Group(children),
	)
}