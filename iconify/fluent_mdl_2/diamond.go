package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1928 829l-904 1113L120 829l537-573h734l537 573zm-666-61l-144-384H930L786 768h476zM779 896l245 858l245-858H779zm-93-512L327 768h343l144-384H686zM314 896l542 667l-191-667H314zm878 667l542-667h-351l-191 667zm529-795l-359-384h-128l144 384h343z"/>`),
		g.Group(children),
	)
}