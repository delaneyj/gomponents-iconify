package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 5a.75.75 0 0 1 .75-.75h11.5a3.75 3.75 0 1 1 0 7.5H9.87l.97.97a.75.75 0 1 1-1.06 1.06l-2.25-2.25L7 11l.53-.53l2.25-2.25a.75.75 0 1 1 1.06 1.06l-.97.97h2.38a2.25 2.25 0 0 0 0-4.5H.75A.75.75 0 0 1 0 5Zm6 6a.75.75 0 0 0-.75-.75H.75a.75.75 0 0 0 0 1.5h4.5A.75.75 0 0 0 6 11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}