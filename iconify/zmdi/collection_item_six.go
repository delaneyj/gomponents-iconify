package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionItemSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91h43zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5h299zm0 342V48H128v299h299zm-171-43q-18 0-30.5-12.5T213 261V133q0-17 12.5-29.5T256 91h85v42h-85v43h43q17 0 29.5 12.5T341 219v42q0 18-12.5 30.5T299 304h-43zm0-85v42h43v-42h-43z"/>`),
		g.Group(children),
	)
}