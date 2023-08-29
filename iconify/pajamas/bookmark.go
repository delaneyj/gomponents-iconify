package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8 9.293l1.06 1.06l3.44 3.44V2a.5.5 0 0 0-.5-.5H4a.5.5 0 0 0-.5.5v11.793l3.44-3.44L8 9.293Zm1.06 3.182l3.233 3.232c.63.63 1.707.184 1.707-.707V2a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v13c0 .89 1.077 1.337 1.707.707l3.232-3.232L8 11.415l1.06 1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}