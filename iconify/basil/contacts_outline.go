package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 9a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm1 3.5a3 3 0 0 0-3 3a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1a3 3 0 0 0-3-3h-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.543 2.883a31.331 31.331 0 0 1 8.913 0a3.196 3.196 0 0 1 2.73 2.874l.126 1.396c.293 3.225.293 6.47 0 9.694l-.127 1.396a3.197 3.197 0 0 1-2.729 2.874a31.334 31.334 0 0 1-8.913 0a3.197 3.197 0 0 1-2.728-2.874l-.127-1.396a53.504 53.504 0 0 1 0-9.694l.127-1.396a3.197 3.197 0 0 1 2.728-2.874Zm8.7 1.484a29.832 29.832 0 0 0-8.486 0a1.697 1.697 0 0 0-1.448 1.526l-.127 1.396a52.003 52.003 0 0 0 0 9.422l.127 1.396c.07.783.67 1.414 1.448 1.526a29.86 29.86 0 0 0 8.486 0a1.696 1.696 0 0 0 1.448-1.526l.127-1.396a52.009 52.009 0 0 0 0-9.422l-.127-1.396a1.697 1.697 0 0 0-1.448-1.526Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}