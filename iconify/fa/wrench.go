package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M384 1344q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45zm644-420l-682 682q-37 37-90 37q-52 0-91-37L59 1498q-38-36-38-90q0-53 38-91l681-681q39 98 114.5 173.5T1028 924zm634-435q0 39-23 106q-47 134-164.5 217.5T1216 896q-185 0-316.5-131.5T768 448t131.5-316.5T1216 0q58 0 121.5 16.5T1445 63q16 11 16 28t-16 28l-293 169v224l193 107q5-3 79-48.5t135.5-81T1630 454q15 0 23.5 10t8.5 25z"/>`),
		g.Group(children),
	)
}