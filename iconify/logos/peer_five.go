package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeerFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path fill="#E44D26" d="m0 0l23.297 261.31l104.546 29.022l104.835-29.064L256 0z"/><path fill="#EF6429" d="m128 268.117l84.712-23.485l19.93-223.266H128z"/><path fill="#EBEBEB" d="M80.798 84.426H128l1.001-1.364V54.505L128 52.412H47.481L60.82 232.368l31.043 9.216l-5.092-78.336H128l1.001-1.425v-28.467L128 131.234H83.678z"/><path fill="#FEFEFE" d="M128 52.412h75.184l-8.003 110.835H128v-32.013h36.652l3.366-46.808H128z"/></g>`),
		g.Group(children),
	)
}