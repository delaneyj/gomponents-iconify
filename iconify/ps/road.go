package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M17 495q14 17 34 17h412q22 0 34-17q14-18 6-38L367 30Q356 0 326 0H186q-13 0-25 8.5T145 30L9 457q-5 20 8 38zM186 43h49v64q0 21 21 21t21-21V43h49l137 426H299v-85q0-21-22-21h-42q-22 0-22 21v85H51zm53 256h34q10 0 16.5-8t4.5-18l-15-85q-3-17-21-17h-6q-16 0-22 17l-15 85q0 10 7 18t17 8z"/>`),
		g.Group(children),
	)
}