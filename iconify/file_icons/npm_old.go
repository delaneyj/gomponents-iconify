package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NpmOld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M261 269h-27v-58h27v58zm59.5 29H379v-87h28.5v87l29 1v-88H467v87h29V182H320.5v116zm-146 29H234v-28.5h57v-58l1-58.5H174.5v145zm-146-28.5h60V211h28v87h30V182h-118v116.5zM524 327H263.5v30.5h-117v-30H0V153h524v174z"/>`),
		g.Group(children),
	)
}