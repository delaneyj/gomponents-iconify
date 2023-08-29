package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStoolCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 14a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Zm2-9.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z"/><path d="M9 9.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V10a.5.5 0 0 1 .5-.5Zm8 0a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V10a.5.5 0 0 1 .5-.5Z"/><path d="M17 4.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Zm1.413 16.782a.5.5 0 0 1-.695.13l-9.5-6.5a.5.5 0 0 1 .564-.825l9.5 6.5a.5.5 0 0 1 .13.695Z"/><path d="M7.587 21.282a.5.5 0 0 0 .695.13l9.5-6.5a.5.5 0 0 0-.564-.825l-9.5 6.5a.5.5 0 0 0-.13.695ZM17 7H9V6h8v1Zm0 3H9V9h8v1ZM5.5 11a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Zm11 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}