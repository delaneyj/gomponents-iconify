package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Currencies(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.5 19.125l-13 11.917l13 11.916m13-37.916l13 13l-13 13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.817 33.967a2.583 2.583 0 0 0 2.166.975h1.3a2.167 2.167 0 0 0 0-4.334h-1.408a2.167 2.167 0 1 1 0-4.333h1.3a2.326 2.326 0 0 1 2.167.975m-2.817-2.167v10.834m18.308-14.625a2.88 2.88 0 0 1-2.275 1.083a2.913 2.913 0 0 1-2.925-2.925v-2.817a2.913 2.913 0 0 1 2.925-2.925a3.049 3.049 0 0 1 2.275 1.084m-6.5 2.058h3.684m-3.684 2.167h3.684"/>`),
		g.Group(children),
	)
}