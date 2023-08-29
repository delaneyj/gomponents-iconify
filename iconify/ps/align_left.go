package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 85h214q21 0 21-21t-21-21H21Q0 43 0 64t21 21zm0 86h299q21 0 21-22q0-21-21-21H21q-9 0-15 6t-6 15q0 10 6 16t15 6zm0 85h256q10 0 16-6t6-15q0-22-22-22H21q-9 0-15 6t-6 16q0 9 6 15t15 6zm0 85h342q21 0 21-21t-21-21H21q-21 0-21 21t21 21z"/>`),
		g.Group(children),
	)
}