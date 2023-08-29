package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firstaid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V384q0-53 37.5-90.5T128 256h128v-64q0-80 56-136T448 0h128q80 0 136 56t56 136v64h128q53 0 90.5 37.5T1024 384v512q0 53-37.5 90.5T896 1024zM640 192q0-27-19-45.5T576 128H448q-27 0-45.5 18.5T384 192v64h256v-64zm64 384H576V448q0-27-19-45.5T511.5 384t-45 18.5T448 448v128H320q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45.5 19h128v128q0 26 18.5 45t45 19t45.5-19t19-45V704h128q27 0 45.5-19t18.5-45.5t-18.5-45T704 576z"/>`),
		g.Group(children),
	)
}