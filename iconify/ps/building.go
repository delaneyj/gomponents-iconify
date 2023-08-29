package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M85 512h278q9 0 15-6t6-15V43q0-18-12.5-30.5T341 0H43Q25 0 12.5 12.5T0 43v448q0 9 6 15t15 6h64zm86-43h-43v-85h43v85zm85 0h-43v-85h43v85zM43 43h298v426h-42v-85q0-17-13-30t-30-13H128q-17 0-30 13t-13 30v85H43V43zm42 42h43v43H85V85zm86 0h42v43h-42V85zm85 0h43v43h-43V85zM85 171h43v42H85v-42zm86 0h42v42h-42v-42zm85 0h43v42h-43v-42zM85 256h43v43H85v-43zm86 0h42v43h-42v-43zm85 0h43v43h-43v-43z"/>`),
		g.Group(children),
	)
}