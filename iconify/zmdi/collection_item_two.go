package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionItemTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91h43zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5h299zm0 342V48H128v299h299zm-86-86v43H213v-85q0-18 12.5-30.5T256 176h43v-43h-86V91h86q17 0 29.5 12.5T341 133v43q0 18-12.5 30.5T299 219h-43v42h85z"/>`),
		g.Group(children),
	)
}