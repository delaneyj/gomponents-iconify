package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodepenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.21.093a.5.5 0 0 1 .58 0l7 5A.5.5 0 0 1 15 5.5v4a.5.5 0 0 1-.21.407l-7 5a.5.5 0 0 1-.58 0l-7-5A.5.5 0 0 1 0 9.5v-4a.5.5 0 0 1 .21-.407l7-5ZM1 6.472L2.44 7.5L1 8.528V6.472ZM1.36 9.5L7 13.528v-2.77L3.3 8.113L1.36 9.5Zm2.8-2L7.5 9.886L10.84 7.5L7.5 5.114L4.16 7.5ZM8 4.243l3.7 2.643L13.64 5.5L8 1.472v2.77ZM7 1.472v2.77L3.3 6.887L1.36 5.5L7 1.472Zm7 5L12.56 7.5L14 8.528V6.472ZM13.64 9.5L11.7 8.114L8 10.757v2.771L13.64 9.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}