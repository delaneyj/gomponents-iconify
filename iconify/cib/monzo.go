package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monzo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.984 23.016H16l7.907-7.896v14.719c.02.521.64.771 1.015.416l6.641-6.629c.276-.281.432-.667.421-1.063V7.052L32 7.041l-5.307-5.296a.587.587 0 0 0-.839 0l-9.855 9.843l-9.88-9.869a.587.587 0 0 0-.839 0L-.001 7.042v15.547c-.005.391.151.772.427 1.052l6.625 6.641a.603.603 0 0 0 1.016-.427l.011-14.751l7.905 7.923z"/>`),
		g.Group(children),
	)
}