package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heavymetal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M700 148q-31 208-43 301q-17 126-17 190v97q0 176-97 265q-17 16-31 23H128q-30 0-61-51.5T18 861T0 767V512q0-26 18.5-45T64 448q19 0 36 12L64 70q-2-28 15-48.5T122 0t46 17.5T189 63q35 321 35 353q1-4 3.5-10.5T239 382t20-30t29-23t38-9q24 1 41 20t17 46v-2q0-26 19-45t45-19t45 19t19 45v-1q5-10 12-22l65-257q9-25 32-36t46-2t32.5 33t.5 49zM384 384z"/>`),
		g.Group(children),
	)
}