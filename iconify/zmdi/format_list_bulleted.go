package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListBulleted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M32 160q13 0 22.5 9.5T64 192t-9.5 22.5T32 224t-22.5-9.5T0 192t9.5-22.5T32 160zm0-128q13 0 22.5 9.5T64 64t-9.5 22.5T32 96T9.5 86.5T0 64t9.5-22.5T32 32zm0 260q12 0 20 8t8 20t-8 20t-20 8t-20-8t-8-20t8-20t20-8zm64 49v-42h299v42H96zm0-128v-42h299v42H96zm0-170h299v42H96V43z"/>`),
		g.Group(children),
	)
}