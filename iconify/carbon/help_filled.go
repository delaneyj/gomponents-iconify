package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm0 23a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 25Zm1.142-7.754v2.501h-2.25V15h2.125a2.376 2.376 0 0 0 0-4.753h-1.5a2.378 2.378 0 0 0-2.375 2.375v.638h-2.25v-.638A4.628 4.628 0 0 1 15.517 8h1.5a4.624 4.624 0 0 1 .125 9.246Z"/><path fill="none" d="M16 25a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 25Zm1.142-7.754v2.501h-2.25V15h2.125a2.376 2.376 0 0 0 0-4.753h-1.5a2.378 2.378 0 0 0-2.375 2.375v.638h-2.25v-.638A4.628 4.628 0 0 1 15.517 8h1.5a4.624 4.624 0 0 1 .125 9.246Z"/>`),
		g.Group(children),
	)
}