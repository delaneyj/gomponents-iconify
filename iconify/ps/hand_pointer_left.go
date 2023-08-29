package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointerLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M237 341h43q68 0 112.5-50.5T427 173q-6-55-50.5-92.5T276 43H152q0 17 12.5 29.5T195 85h81q42 0 74 27.5t37 68.5q4 48-28 83t-79 35h-85q0 17 12.5 29.5T237 341zm-85-64h64q21 0 21-21t-21-21h-64q-21 0-21 21t21 21zm0-64h64q21 0 21-21t-21-21h-64q-21 0-21 21t21 21zM3 128q0 21 21 21h192v-42H24q-21 0-21 21z"/>`),
		g.Group(children),
	)
}