package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaguetteBread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.861 5.14c-2.853-2.854-7.477-2.852-10.336-.002L5.139 16.524a7.314 7.314 0 0 0 0 10.337a7.314 7.314 0 0 0 10.337 0l11.385-11.385a7.314 7.314 0 0 0 0-10.337Zm-8.924 1.414a5.314 5.314 0 0 1 7.51 0a5.314 5.314 0 0 1 0 7.508L14.062 25.447a5.314 5.314 0 0 1-7.509 0a5.314 5.314 0 0 1 0-7.51l.128-.127l3.732 3.731a1.556 1.556 0 0 0 2.206 0a1.556 1.556 0 0 0 0-2.206l-3.732-3.731l2.255-2.255l3.736 3.736a1.556 1.556 0 0 0 2.207 0a1.572 1.572 0 0 0 0-2.207l-3.736-3.736l2.255-2.255l3.731 3.732a1.556 1.556 0 0 0 2.207 0a1.556 1.556 0 0 0 0-2.207l-3.732-3.73l.127-.128Z"/>`),
		g.Group(children),
	)
}