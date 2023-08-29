package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.5 11A1.5 1.5 0 0 1 5 9.5h10a1.5 1.5 0 0 1 0 3H5A1.5 1.5 0 0 1 3.5 11ZM6 1a1 1 0 0 1 1 1v5a1 1 0 0 1-2 0V2a1 1 0 0 1 1-1Z"/><path d="M6 6a1 1 0 0 1 1 1v2.5a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1Zm8 0a1 1 0 0 1 1 1v2.5a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1Z"/><path d="M14 1a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0V2a1 1 0 0 1 1-1Zm1.825 17.565a1 1 0 0 1-1.39.26l-9.5-6.5a1 1 0 0 1 1.13-1.65l9.5 6.5a1 1 0 0 1 .26 1.39Z"/><path d="M4.175 18.565a1 1 0 0 0 1.39.26l9.5-6.5a1 1 0 1 0-1.13-1.65l-9.5 6.5a1 1 0 0 0-.26 1.39ZM14 4H6V2h8v2Zm0 3H6V5h8v2ZM2 8a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H3a1 1 0 0 1-1-1Zm12 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}