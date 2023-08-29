package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M67 512h362v-43H67q-22 0-22-21t22-21h362V0H67Q39 0 21 18.5T3 64v363h4q-4 12-4 21q0 27 18 45.5T67 512zM45 64q0-21 22-21h320v341H67q-8 0-22 4V64zm86 107h192v42H131v-42zm0-86h192v43H131V85z"/>`),
		g.Group(children),
	)
}