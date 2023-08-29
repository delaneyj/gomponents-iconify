package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M363 85q21 0 21-21t-21-21H149q-21 0-21 21t21 21h214zM43 149q0 22 21 22h299q9 0 15-6t6-16q0-9-6-15t-15-6H64q-21 0-21 21zm320 64H107q-22 0-22 22q0 9 6 15t16 6h256q9 0 15-6t6-15q0-10-6-16t-15-6zM0 320q0 21 21 21h342q21 0 21-21t-21-21H21q-21 0-21 21z"/>`),
		g.Group(children),
	)
}