package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M365 288H67q-28 0-46 18.5T3 352v43q0 27 18 45.5T67 459h298q28 0 46-18.5t18-45.5v-43q0-27-18-45.5T365 288zm22 107q0 9-6 15t-16 6H67q-10 0-16-6t-6-15v-43q0-21 22-21h298q22 0 22 21v43zM254 15Q241 0 218 0q-25 0-36 15L18 192q-25 22-11 49q10 26 38 26h342q26 0 40-26q11-28-8-49zm133 209H45v-4L210 43q2-1 6-1t4 1z"/>`),
		g.Group(children),
	)
}