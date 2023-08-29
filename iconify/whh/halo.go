package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Halo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M570 384h38q13 0 22.5 9.5T640 416t-9.5 22.5T608 448h-38q-10 28-34.5 46T480 512q-40 0-68-28t-28-68t28-68t68-28q31 0 55.5 18t34.5 46zm422 64h-34q-9 101-71.5 186.5T725 772t-213 59v33q0 13-9.5 22.5T480 896t-22.5-9.5T448 864v-98q-106-8-192-59v-76q97 73 224 73q74 0 137-15.5t111.5-46t76-80.5T832 448q0-87-47-160.5T657 171t-177-43t-177 43t-128 116.5T128 448H32q-13 0-22.5-9.5T0 416t9.5-22.5T32 384h38q18-98 85-175.5T320 94V41Q412 0 512 0q116 0 215.5 51t161 139T958 384h34q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 448z"/>`),
		g.Group(children),
	)
}