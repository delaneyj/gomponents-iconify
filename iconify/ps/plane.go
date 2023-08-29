package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m489 230l-171-74l2-49q0-34-21.5-70.5T245 0t-53 36.5t-21 72.5l2 49l-150 75Q0 243 0 271v77l181-41l5 86q-79 32-79 98v21h277v-21q0-71-79-101l4-89l203 45v-77q0-24-23-39zM43 294v-23l132-66l4 60zm111 175q5-12 34-32l2 32h-36zm79 0l-20-362q0-19 12-41.5T245 43t20 22t12 40l-19 364h-25zm104 0h-38l2-34q29 13 36 34zm132-175l-158-34l3-57l153 68v23h2z"/>`),
		g.Group(children),
	)
}