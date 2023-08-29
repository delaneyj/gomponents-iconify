package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionItemEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91h43zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5h299zm0 342V48H128v299h299zm-171-43q-18 0-30.5-12.5T213 261v-32q0-13 9.5-22.5T245 197q-13 0-22.5-9t-9.5-23v-32q0-17 12.5-29.5T256 91h43q17 0 29.5 12.5T341 133v32q0 14-9 23t-23 9q14 0 23 9.5t9 22.5v32q0 18-12.5 30.5T299 304h-43zm0-171v43h43v-43h-43zm0 86v42h43v-42h-43z"/>`),
		g.Group(children),
	)
}