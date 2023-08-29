package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketLaunch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.808 0a9.999 9.999 0 0 0-7.142 3H6.75A2.75 2.75 0 0 0 4 5.75V8h2l2 2v2h2.25A2.75 2.75 0 0 0 13 9.25V7.334a9.998 9.998 0 0 0 3-7.142V0h-.192ZM6.44 6.5a9.964 9.964 0 0 1 1.015-2H6.75c-.69 0-1.25.56-1.25 1.25v.75h.94Zm3.06 4v-.94a9.966 9.966 0 0 0 2-1.015v.705c0 .69-.56 1.25-1.25 1.25H9.5Zm4.88-8.88a8.502 8.502 0 0 0-6.71 5.928l.782.783a8.502 8.502 0 0 0 5.928-6.71Zm-11.6 8.66a.75.75 0 1 0-1.06-1.06l-1.5 1.5a.75.75 0 1 0 1.06 1.06l1.5-1.5Zm3 1a.75.75 0 1 0-1.06-1.06l-4.5 4.5a.75.75 0 1 0 1.06 1.06l4.5-4.5Zm1 1.94a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 0 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}