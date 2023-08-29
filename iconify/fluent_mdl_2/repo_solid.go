package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 768h640v1152H128V768h640v768l256-128l256 128V768zm768-640v512H0V128h2048z"/>`),
		g.Group(children),
	)
}