package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CdOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCd1x10"><path fill="#fff" d="M0-88h600v600H0z"/></clipPath></defs><g clip-path="url(#flagCd1x10)" transform="matrix(.853 0 0 .853 0 75.1)"><path fill="#007fff" d="M0-88h800v600H0z"/><path fill="#f7d618" d="M36 32h84l26-84l26 84h84l-68 52l26 84l-68-52l-68 52l26-84l-68-52zM750-88L0 362v150h50L800 62V-88h-50"/><path fill="#ce1021" d="M800-88L0 392v120L800 32V-88"/></g>`),
		g.Group(children),
	)
}