package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnowyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.027 17.43A8.003 8.003 0 0 1 9 2a8.003 8.003 0 0 1 7.458 5.099A5.5 5.5 0 1 1 18 17.978a6 6 0 0 0-11.973-.549ZM13 16.267l1.964-1.134l1 1.732L14 18l1.964 1.134l-1 1.732L13 19.732V22h-2v-2.268l-1.964 1.134l-1-1.732L10 18l-1.964-1.134l1-1.732L11 16.268V14h2v2.268Z"/>`),
		g.Group(children),
	)
}