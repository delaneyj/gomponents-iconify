package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilReply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 336q0 66-25 127t-73 110L633 1890L0 2048l158-633L1475 98q48-48 109-73t128-25q70 0 131 26t106 72t72 107t27 131zm-128 0q0-43-16-81t-45-66t-66-44t-81-17q-38 0-66 10t-53 29t-47 41t-47 48l293 293q25-25 48-47t41-46t28-53t11-67zM326 1428q107 35 183 111t111 183L1701 640l-293-293L326 1428zm-150 444l329-82q-10-46-32-87t-55-73t-73-54t-87-33l-82 329zm1616-336q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20v-128q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10h-293l162 163l-90 90l-318-317l318-317l90 90l-162 163h293z"/>`),
		g.Group(children),
	)
}