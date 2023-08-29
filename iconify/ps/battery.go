package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 107q0-28-18.5-46T384 43H64q-27 0-45.5 18T0 107v170q0 28 18.5 46T64 341h320q27 0 45.5-18t18.5-46q27 0 45.5-18t18.5-46v-42q0-28-18.5-46T448 107zm21 106q0 22-21 22q-17 0-30 12.5T405 277q0 22-21 22H64q-21 0-21-22V107q0-22 21-22h320q21 0 21 22q0 17 13 29.5t30 12.5q21 0 21 22v42zM85 128h235v128H85V128z"/>`),
		g.Group(children),
	)
}