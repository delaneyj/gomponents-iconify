package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoentrepreneur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M23.526 14.073a8.573 8.573 0 1 1-17.146 0a8.573 8.573 0 0 1 17.146 0Z"/><path fill="none" stroke="currentColor" d="M42.477 24a18.5 18.5 0 0 1-36.952 1.32c-.034-.463.347-.869.855-.869h16.244c1.004 0 1.783-.81 1.852-1.771A18.497 18.497 0 0 1 35.849 6.908a18.5 18.5 0 0 1 5.76-1.361c.462-.033.869.348.869.855V24Z"/>`),
		g.Group(children),
	)
}