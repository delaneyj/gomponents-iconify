package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCentered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M85 43q-21 0-21 21t21 21h214q21 0 21-21t-21-21H85zM21 149q0 22 22 22h298q22 0 22-22q0-9-6-15t-16-6H43q-10 0-16 6t-6 15zm86 64q-22 0-22 22q0 9 6 15t16 6h170q10 0 16-6t6-15q0-22-22-22H107zM0 320q0 21 21 21h342q21 0 21-21t-21-21H21q-21 0-21 21z"/>`),
		g.Group(children),
	)
}