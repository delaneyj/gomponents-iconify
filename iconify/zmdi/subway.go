package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M170.5 3q39.5 0 67 3t53 11.5t38 26T341 88v203q0 31-21.5 52.5T267 365l32 32v11H43v-11l32-32q-31 0-53-21.5T0 291V88q0-27 12.5-44.5t38-26t53-11.5t67-3zm-96 320q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zM149 195V88H43v107h106zm117.5 128q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zM299 195V88H192v107h107z"/>`),
		g.Group(children),
	)
}