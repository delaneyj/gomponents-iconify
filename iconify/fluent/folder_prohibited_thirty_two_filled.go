package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderProhibitedThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8.5A4.5 4.5 0 0 1 6.5 4h3.964a3.5 3.5 0 0 1 2.475 1.025L15 7.085l-3.475 3.476a1.5 1.5 0 0 1-1.06.439H2V8.5ZM2 13v10.5A4.5 4.5 0 0 0 6.5 28h10.015A9 9 0 0 1 30 16.292V12a4.5 4.5 0 0 0-4.5-4.5h-8.086l-4.475 4.475A3.5 3.5 0 0 1 10.464 13H2Zm22 17.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm0-2a5.475 5.475 0 0 1-3.117-.968l7.649-7.65A5.5 5.5 0 0 1 24 28.5Zm3.118-10.032l-7.65 7.65a5.5 5.5 0 0 1 7.65-7.65Z"/>`),
		g.Group(children),
	)
}