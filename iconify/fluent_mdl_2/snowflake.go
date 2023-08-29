package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1871 1276l33 124l-200 53l53 201l-124 33l-63-237l-546-315v631l173 173l-90 90l-147-146l-147 146l-90-90l173-173v-631l-546 315l-63 237l-124-33l53-201l-200-53l33-124l237 63l546-315l-546-315l-237 63l-33-124l200-53l-53-201l124-33l63 237l546 315V282L723 109l90-90l147 146l147-146l90 90l-173 173v631l546-315l63-237l124 33l-53 201l200 53l-33 124l-237-63l-546 315l546 315l237-63z"/>`),
		g.Group(children),
	)
}