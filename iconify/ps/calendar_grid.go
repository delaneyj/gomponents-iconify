package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 43h-43V21q0-21-21-21t-21 21v22H149V21q0-21-21-21t-21 21v22H64q-27 0-45.5 19.5T0 109v358q0 19 12.5 32T43 512h426q18 0 30.5-13t12.5-32V109q0-27-18.5-46.5T448 43zM107 469H43v-42h64v42zm0-85H43v-43h64v43zm0-85H43v-43h64v43zm0-86H43v-42h64v42zm85 256h-43v-42h43v42zm0-85h-43v-43h43v43zm0-85h-43v-43h43v43zm0-86h-43v-42h43v42zm85 256h-42v-42h42v42zm0-85h-42v-43h42v43zm0-85h-42v-43h42v43zm0-86h-42v-42h42v42zm86 256h-43v-42h43v42zm0-85h-43v-43h43v43zm0-85h-43v-43h43v43zm0-86h-43v-42h43v42zM43 128v-19q0-10 6-17t15-7h384q21 0 21 22v21H43zm426 43v42h-64v-42h64zm-64 170h64v43h-64v-43zm64 128h-64v-42h64v42zm0-170h-64v-43h64v43z"/>`),
		g.Group(children),
	)
}