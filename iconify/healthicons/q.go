package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Q(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 24c0-7.732 6.268-14 14-14s14 6.268 14 14c0 3.145-1.037 6.047-2.787 8.384l2.201 2.202a2 2 0 1 1-2.828 2.828l-2.202-2.201A13.938 13.938 0 0 1 24 38c-7.732 0-14-6.268-14-14Zm22.343 5.515l-2.929-2.93a2 2 0 1 0-2.828 2.83l2.929 2.928A9.954 9.954 0 0 1 24 34c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10c0 2.038-.61 3.934-1.657 5.515Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}