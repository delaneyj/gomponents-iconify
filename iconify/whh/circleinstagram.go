package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleinstagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm256-704q0-26-18.5-45T704 256H320q-26 0-45 19t-19 45v64h114q58-64 142-64t142 64h114v-64zM384 512q0 53 37.5 90.5T512 640t90.5-37.5T640 512t-37.5-90.5T512 384t-90.5 37.5T384 512zm384-64h-76q12 33 12 64q0 80-56 136t-136 56t-136-56t-56-136q0-31 12-64h-76v256q0 27 18.5 45.5T320 768h384q27 0 45.5-18.5T768 704V448z"/>`),
		g.Group(children),
	)
}