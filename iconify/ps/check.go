package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405 491V192l-21-43l-21 43v299q0 21 21 21t21-21zM491 0H21Q12 0 6 6T0 21v320h320v-42H43V107h426v192h-21v42h64V21q0-9-6-15t-15-6zm-22 64H43V43h426v21zM85 149h235v43H85v-43zm0 64h214v43H85v-43z"/>`),
		g.Group(children),
	)
}