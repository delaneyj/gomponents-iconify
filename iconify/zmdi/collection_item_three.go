package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionItemThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5h299zm0 342V48H128v299h299zM43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91h43zm298 170q0 18-12.5 30.5T299 304h-86v-43h86v-42h-43v-43h43v-43h-86V91h86q17 0 29.5 12.5T341 133v32q0 14-9 23t-23 9q14 0 23 9.5t9 22.5v32z"/>`),
		g.Group(children),
	)
}