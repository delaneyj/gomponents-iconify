package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bvnc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 34.673h-5.968s-1.768.07-.864-1.496s2.262.474 2.262-5.07M24 43.5H5.686l3.237-6.752h15.076M24 34.673h5.968s1.768.07.864-1.496s-2.262.474-2.262-5.07M24 43.5h18.314l-3.237-6.752H24.001M6.725 4.5h34.551v23.577H6.725z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 21.177h-7.862v-7.732h2.988V11.72h2.36v-1.016H24m0 10.473h7.862v-7.732h-2.988V11.72h-2.36v-1.016H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M24 21.177v-2.743m-2.296 2.743v-2.743m-2.295 2.743v-2.743m9.182 2.743v-2.743m-2.295 2.743v-2.743"/>`),
		g.Group(children),
	)
}