package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintbucketSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A2.5 2.5 0 0 0 2 2.5v3.793l-.44.44a2.5 2.5 0 0 0 0 3.535l3.172 3.171a2.5 2.5 0 0 0 3.536 0l4.171-4.171a2.5 2.5 0 0 0 0-3.536L9.268 2.561a2.498 2.498 0 0 0-2.342-.666A2.501 2.501 0 0 0 4.5 0ZM6 3.707V7h1V2.914a1.5 1.5 0 0 1 1.56.354l3.172 3.171a1.5 1.5 0 0 1 0 2.122l-.44.439H1.915a1.5 1.5 0 0 1 .354-1.56L6 3.706Zm-.009-1.372A1.5 1.5 0 0 0 3 2.5v2.793L5.732 2.56c.082-.083.169-.158.259-.226Z" clip-rule="evenodd"/><path fill="currentColor" d="m12.645 9.737l1.534 1.534a2.17 2.17 0 1 1-3.069 0l1.535-1.534Z"/>`),
		g.Group(children),
	)
}