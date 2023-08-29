package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m480 147l-60-15l-55-89q-15-22-36-22H43q-17 0-30 13T0 64v256h26q6 17 23 30t36 13q21 0 37.5-12t22.5-31h198q15 43 60 43q20 0 36.5-13t23.5-30h6q18 0 30.5-12.5T512 277v-89q0-14-9-26t-23-15zM85 320q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6zm320 0q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6zm64-43h-4q-6-17-23-29.5T405 235t-36.5 11.5T346 277H145q-6-18-23-30t-37-12q-27 0-42 17V64h286l12 21h-21q-21 0-21 22v42q0 22 21 22h83q1 0 3 1t4 1l59 15v89z"/>`),
		g.Group(children),
	)
}