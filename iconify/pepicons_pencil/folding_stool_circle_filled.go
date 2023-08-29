package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStoolCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFoldingStoolCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 14a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Zm2-9.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z"/><path d="M9 9.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V10a.5.5 0 0 1 .5-.5Zm8 0a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V10a.5.5 0 0 1 .5-.5Z"/><path d="M17 4.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Zm1.413 16.782a.5.5 0 0 1-.695.13l-9.5-6.5a.5.5 0 0 1 .564-.825l9.5 6.5a.5.5 0 0 1 .13.695Z"/><path d="M7.587 21.282a.5.5 0 0 0 .695.13l9.5-6.5a.5.5 0 0 0-.564-.825l-9.5 6.5a.5.5 0 0 0-.13.695ZM17 7H9V6h8v1Zm0 3H9V9h8v1ZM5.5 11a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Zm11 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFoldingStoolCircleFilled0)"/></g>`),
		g.Group(children),
	)
}