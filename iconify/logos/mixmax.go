package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mixmax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#E55CFF" d="M0 202.999h50.731V.007H0v202.992Z"/><path fill="#24235C" d="M102.633.007L0 202.999h50.731L153.364.007h-50.73Z"/><path fill="#E55CFF" d="M102.633 202.999h50.731V.007h-50.731v202.992Z"/><path fill="#24235C" d="M205.27.008L102.635 203h50.73L256 .008h-50.73Z"/><path fill="#E55CFF" d="M204.982 202.999h50.731V.007h-50.731v202.992Z"/>`),
		g.Group(children),
	)
}