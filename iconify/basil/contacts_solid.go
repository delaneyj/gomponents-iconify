package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.456 2.883a31.332 31.332 0 0 0-8.912 0a3.197 3.197 0 0 0-2.73 2.874l-.126 1.396a53.504 53.504 0 0 0 0 9.694l.127 1.396a3.197 3.197 0 0 0 2.729 2.874c2.955.425 5.957.425 8.912 0a3.197 3.197 0 0 0 2.73-2.874l.126-1.396c.293-3.225.293-6.47 0-9.694l-.127-1.396a3.196 3.196 0 0 0-2.729-2.874ZM10 9a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-2 6.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}