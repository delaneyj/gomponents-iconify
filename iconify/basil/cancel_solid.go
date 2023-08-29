package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.345 4.017a42.253 42.253 0 0 1 9.31 0c1.713.192 3.095 1.541 3.296 3.26a40.66 40.66 0 0 1 0 9.445a3.734 3.734 0 0 1-3.296 3.26a42.123 42.123 0 0 1-9.31 0a3.734 3.734 0 0 1-3.296-3.26a40.652 40.652 0 0 1 0-9.444a3.734 3.734 0 0 1 3.295-3.26ZM8.47 8.47a.75.75 0 0 1 1.06 0L12 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L13.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L12 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L10.94 12L8.47 9.53a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}