package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 768q0 106-75 181t-181 75V512q31 0 64 8v-72q0-148-82.5-234T512 128t-237.5 86T192 448v72q33-8 64-8v512q-106 0-181-75T0 768q0-96 64-169V448q0-91 35.5-174T195 131t143-95.5T512 0t174 35.5T829 131t95.5 143T960 448v151q64 73 64 169z"/>`),
		g.Group(children),
	)
}