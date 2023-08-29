package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleTranslate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h9L32.3 5.5Zm9 37h24.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95H32.3M30 22.35l10.11 2.39m-3.85-4.53l.1 3.65M24.71 40c2.65-4 5.67-7.52 13.61-10.67m-9.81-2.93A34.93 34.93 0 0 1 37 38M14.78 15h4.87a4.89 4.89 0 0 1-4.87 5a5 5 0 1 1 2.48-9.36"/>`),
		g.Group(children),
	)
}