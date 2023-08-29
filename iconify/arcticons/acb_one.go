package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcbOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.415 18.071c1.546 0 2.8 1.274 2.8 2.845s-1.254 2.846-2.8 2.846h-4.62V12.38h4.62c1.546 0 2.8 1.274 2.8 2.846s-1.254 2.845-2.8 2.845Zm0 0h-4.621m-3.369 1.874v.047c0 2.082-1.661 3.77-3.71 3.77s-3.71-1.688-3.71-3.77V16.15c0-2.082 1.66-3.77 3.71-3.77h0c2.049 0 3.71 1.688 3.71 3.77v.047m-11.463 3.758h-4.944M9.785 23.73l3.71-11.349l3.709 11.383m12.543 8.103h2.447M33.5 35.62h-3.753v-7.506H33.5M16.986 35.62a2.486 2.486 0 0 1-2.486-2.486V30.6a2.486 2.486 0 1 1 4.972 0v2.533a2.486 2.486 0 0 1-2.486 2.486Zm4.563-7.506v7.506l1.876-1.876m4.222 1.876v-7.506l-1.876 1.877"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}