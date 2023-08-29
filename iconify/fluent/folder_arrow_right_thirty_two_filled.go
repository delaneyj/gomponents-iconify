package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderArrowRightThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8.5A4.5 4.5 0 0 1 6.5 4h3.964a3.5 3.5 0 0 1 2.475 1.025L15 7.085l-3.475 3.476a1.5 1.5 0 0 1-1.06.439H2V8.5ZM2 13v10.5A4.5 4.5 0 0 0 6.5 28h10.015A9 9 0 0 1 30 16.292V12a4.5 4.5 0 0 0-4.5-4.5h-8.086l-4.475 4.475A3.5 3.5 0 0 1 10.464 13H2Zm29.5 10a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0Zm-12.25-1h7.69l-2.72-2.72a.75.75 0 1 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 1 1-1.06-1.06l2.72-2.72h-7.69a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}