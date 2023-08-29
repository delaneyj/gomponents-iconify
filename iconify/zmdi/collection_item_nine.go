package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionItemNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91h43zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5h299zm0 342V48H128v299h299zM299 91q17 0 29.5 12.5T341 133v128q0 18-12.5 30.5T299 304h-86v-43h86v-42h-43q-18 0-30.5-12.5T213 176v-43q0-17 12.5-29.5T256 91h43zm0 85v-43h-43v43h43z"/>`),
		g.Group(children),
	)
}