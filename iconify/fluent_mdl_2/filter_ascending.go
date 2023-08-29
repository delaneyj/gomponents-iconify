package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterAscending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1411 1347l317-317l317 317l-90 90l-163-163v646h-128v-646l-163 163l-90-90zM0 128h2048v219l-768 768v805H768v-805L0 347V128zm1920 165v-37H128v37l768 768v731h256v-731l768-768z"/>`),
		g.Group(children),
	)
}