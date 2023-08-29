package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentBookTwoLibraryContentBooksBookShelfStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3.5" height="13" x=".55" y=".5" rx=".5"/><rect width="3.5" height="11" x="4.05" y="2.5" rx=".5"/><rect width="3" height="11" x="9.26" y="2.3" rx=".5" transform="rotate(-14.05 10.779 7.795)"/><path d="M.55 10h3.5m0-1h3.5m2.5 2l2.88-.72"/></g>`),
		g.Group(children),
	)
}