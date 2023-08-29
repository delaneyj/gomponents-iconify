package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 171h-36L316 87q4-16 4-23q0-27-18.5-45.5T256 0t-45.5 18.5T192 64q0 7 4 23L79 171H43q-18 0-30.5 12.5T0 213v256q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V213q0-17-12.5-29.5T469 171zM256 43q21 0 21 21t-21 21t-21-21t21-21zm-32 76q16 9 32 9t32-9l73 52H151zM43 469V213h426v256H43zm362-192H107q-22 0-22 22q0 9 6 15t16 6h298q10 0 16-6t6-15q0-22-22-22zm-64 86H171q-22 0-22 21t22 21h170q22 0 22-21t-22-21z"/>`),
		g.Group(children),
	)
}