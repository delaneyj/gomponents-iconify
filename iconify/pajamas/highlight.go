package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 14.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75Zm1.5-3.252V9.112l1.29-1.258L6.6 9.688l-1.268 1.31H3.5ZM7.645 8.61l4.715-4.866a.5.5 0 0 0-.028-.722l-1.01-.895l.995-1.123l1.01.895a2 2 0 0 1 .11 2.889l-7.47 7.71H2V8.48l7.594-7.41a2 2 0 0 1 2.723-.066l-.995 1.123a.5.5 0 0 0-.68.016L5.863 6.806L7.645 8.61Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}