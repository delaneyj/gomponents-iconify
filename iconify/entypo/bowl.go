package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.949 7.472c-2.176 2.902-4.095 3.002-7.046 3.152h-.101c-3.591-.002-6.138-1.336-6.138-1.832c-.002-.471 2.298-1.697 5.605-1.819l.59-1.473l-.057-.002c-4.908 0-7.791 1.562-7.791 3.051v2c0 .918.582 8.949 7.582 8.949s8-8.031 8-8.949v-2c0-.391-.201-.787-.584-1.158l-.06.081zm.64-4.77a1 1 0 0 0-1.399.201l-3.608 4.809l2.336-5.838a1 1 0 1 0-1.857-.742L9.802 9.274c2.882-.147 4.277-.227 6.067-2.611l1.919-2.561a1 1 0 0 0-.199-1.4z"/>`),
		g.Group(children),
	)
}