package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 562v796l-562 562H562L0 1358V562L562 0h796l562 562zm-128 53l-487-487H615L128 615v690l487 487h690l487-487V615z"/>`),
		g.Group(children),
	)
}