package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionItemNinePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91h43zm234 149q0 18-12.5 30.5T235 283h-64v-43h64v-21h-22q-17 0-29.5-12.5T171 176v-21q0-18 12.5-30.5T213 112h22q17 0 29.5 12.5T277 155v85zm-64-64h22v-21h-22v21zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5h299zm0 171V48H128v299h299V219h-43v42h-43v-42h-42v-43h42v-43h43v43h43z"/>`),
		g.Group(children),
	)
}