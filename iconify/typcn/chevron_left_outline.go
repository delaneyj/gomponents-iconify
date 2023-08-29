package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20a2.98 2.98 0 0 1-2.122-.879L3.757 12l7.122-7.121c1.133-1.133 3.11-1.133 4.243 0C15.688 5.445 16 6.199 16 7s-.312 1.555-.879 2.122L12.243 12l2.878 2.879c.567.566.879 1.32.879 2.121s-.312 1.555-.879 2.122A2.98 2.98 0 0 1 13 20zm-6.415-8l5.708 5.707a1.024 1.024 0 0 0 1.414 0c.189-.189.293-.439.293-.707s-.104-.518-.293-.707L9.415 12l4.292-4.293c.189-.189.293-.44.293-.707s-.104-.518-.293-.707a1.023 1.023 0 0 0-1.414-.001L6.585 12z"/>`),
		g.Group(children),
	)
}