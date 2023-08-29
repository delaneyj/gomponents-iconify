package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStoolCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFoldingStoolCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.5 14A1.5 1.5 0 0 1 8 12.5h10a1.5 1.5 0 0 1 0 3H8A1.5 1.5 0 0 1 6.5 14ZM9 4a1 1 0 0 1 1 1v5a1 1 0 0 1-2 0V5a1 1 0 0 1 1-1Z"/><path d="M9 9a1 1 0 0 1 1 1v2.5a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Zm8 0a1 1 0 0 1 1 1v2.5a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Z"/><path d="M17 4a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0V5a1 1 0 0 1 1-1Zm1.825 17.565a1 1 0 0 1-1.39.26l-9.5-6.5a1 1 0 0 1 1.13-1.65l9.5 6.5a1 1 0 0 1 .26 1.39Z"/><path d="M7.175 21.565a1 1 0 0 0 1.39.26l9.5-6.5a1 1 0 1 0-1.13-1.65l-9.5 6.5a1 1 0 0 0-.26 1.39ZM17 7H9V5h8v2Zm0 3H9V8h8v2ZM5 11a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1Zm12 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFoldingStoolCircleFilled0)"/></g>`),
		g.Group(children),
	)
}