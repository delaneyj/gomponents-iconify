package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M431 107q25-1 42.5-19.5T491 43V0H107v21H21q-9 0-15 6T0 43q0 9 6 15t15 6h22v107q0 43 23.5 80t61.5 54l-21 122l-22 85h427l-21-85zm13 298H154l59-298h171zm4-362q0 9-6 15t-15 6H171q-10 0-16-6t-6-15h299zM85 171V64h26q6 17 21 29.5t34 13.5l-29 155q-52-31-52-91zm52 298l6-21h311l7 21H137z"/>`),
		g.Group(children),
	)
}