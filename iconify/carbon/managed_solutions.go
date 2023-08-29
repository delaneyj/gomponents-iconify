package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManagedSolutions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 23h-5v-5h-2v5h-5v2h5v5h2v-5h5v-2z"/><path fill="currentColor" d="m24.127 11.84l1.181.213a5.792 5.792 0 0 1 2.625 1.144A5.422 5.422 0 0 1 29.953 18h2.03a7.502 7.502 0 0 0-6.15-7.885a10.007 10.007 0 0 0-7.94-7.933a10.002 10.002 0 0 0-11.72 7.933A7.505 7.505 0 0 0 .059 18.41A7.684 7.684 0 0 0 7.773 25H14v-2H7.698a5.632 5.632 0 0 1-5.603-4.486a5.506 5.506 0 0 1 4.434-6.43l1.349-.245l.214-1.11a8.206 8.206 0 0 1 6.742-6.642a7.971 7.971 0 0 1 3.014.13a8.144 8.144 0 0 1 6.053 6.446Z"/>`),
		g.Group(children),
	)
}