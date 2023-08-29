package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStoolCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 14A1.5 1.5 0 0 1 8 12.5h10a1.5 1.5 0 0 1 0 3H8A1.5 1.5 0 0 1 6.5 14ZM9 4a1 1 0 0 1 1 1v5a1 1 0 0 1-2 0V5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9 9a1 1 0 0 1 1 1v2.5a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Zm8 0a1 1 0 0 1 1 1v2.5a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 4a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0V5a1 1 0 0 1 1-1Zm1.825 17.565a1 1 0 0 1-1.39.26l-9.5-6.5a1 1 0 0 1 1.13-1.65l9.5 6.5a1 1 0 0 1 .26 1.39Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.175 21.565a1 1 0 0 0 1.39.26l9.5-6.5a1 1 0 1 0-1.13-1.65l-9.5 6.5a1 1 0 0 0-.26 1.39ZM17 7H9V5h8v2Zm0 3H9V8h8v2ZM5 11a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1Zm12 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}