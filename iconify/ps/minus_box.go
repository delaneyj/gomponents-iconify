package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405 0H107Q62 0 31 31T0 107v298q0 45 31 76t76 31h298q45 0 76-31t31-76V107q0-45-31-76T405 0zm64 405q0 28-18 46t-46 18H107q-28 0-46-18t-18-46V107q0-28 18-46t46-18h298q28 0 46 18t18 46v298zM363 235H149q-21 0-21 21t21 21h214q21 0 21-21t-21-21z"/>`),
		g.Group(children),
	)
}