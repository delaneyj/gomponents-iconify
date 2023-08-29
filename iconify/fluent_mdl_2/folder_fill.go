package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 64q26 0 45 19t19 45v992q0 57 20 93t44 67t44 67t20 93v480q0 26-19 45t-45 19H320V64h1344z"/>`),
		g.Group(children),
	)
}