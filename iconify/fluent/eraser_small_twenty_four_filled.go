package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserSmallTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.651 2.638a2.25 2.25 0 0 1 3.19.023l4.91 5.008a2.25 2.25 0 0 1-.007 3.158l-2.257 2.281a4.5 4.5 0 0 0-5.368 5.424l-1.784 1.803a2.25 2.25 0 0 1-3.176.023l-5.1-5.008a2.25 2.25 0 0 1-.006-3.204l9.598-9.508ZM4.108 13.212a.75.75 0 0 0 .003 1.068l5.099 5.008a.75.75 0 0 0 1.058-.008l1.468-1.483l-6.099-6.1l-1.529 1.515ZM17.5 21a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"/>`),
		g.Group(children),
	)
}