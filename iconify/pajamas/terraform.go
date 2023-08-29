package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.321 7.576L1 5.051V0l4.321 2.525v5.051ZM10.117 16l-4.321-2.525V8.424l4.321 2.525V16ZM5.796 2.819l4.321 2.528v5.048L5.796 7.869V2.82Zm4.796 7.576l4.321-2.523V2.819l-4.321 2.528v5.048Z"/>`),
		g.Group(children),
	)
}