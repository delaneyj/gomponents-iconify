package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lampalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 576H576v320h160q13 0 22.5 9.5T768 928v32h32q13 0 22.5 9.5T832 992t-9.5 22.5t-22.5 9.5H224q-13 0-22.5-9.5T192 992t9.5-22.5T224 960h32v-32q0-13 9.5-22.5T288 896h160V576H256v160q0 13-9.5 22.5T224 768t-22.5-9.5T192 736V576H32q-13 0-22.5-9.5T0 544q0-120 59-224.5T220.5 149T448 68v-4q0-27 19-45.5T512.5 0t45 18.5T576 64v4q125 15 227.5 81T965 319.5t59 224.5q0 13-9.5 22.5T992 576zM416 128q-30 0-67 13t-76 40.5t-71.5 64.5t-53 90T128 448q5-10 14.5-27.5T183 359t63-78t78.5-61.5T416 192q13 0 22.5-9.5T448 160t-9.5-22.5T416 128z"/>`),
		g.Group(children),
	)
}