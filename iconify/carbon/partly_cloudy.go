package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartlyCloudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 15h4v2h-4zm-4-7.413l3-3L27.415 6l-3 3zM15 1h2v4h-2zM4.586 26l3-3l1.415 1.415l-3 3zM4.585 6L6 4.587l3 3l-1.414 1.415z"/><path fill="currentColor" d="M1 15h4v2H1zm25.794 5.342a6.962 6.962 0 0 0-1.868-3.267A8.485 8.485 0 0 0 25 16a9 9 0 1 0-14.585 7.033A4.977 4.977 0 0 0 15 30h10a4.995 4.995 0 0 0 1.794-9.658zM9 16a6.995 6.995 0 0 1 13.985-.297A6.888 6.888 0 0 0 20 15a7.04 7.04 0 0 0-6.794 5.342a4.986 4.986 0 0 0-1.644 1.048A6.968 6.968 0 0 1 9 16zm16 12H15a2.994 2.994 0 0 1-.696-5.908l.658-.157l.099-.67a4.992 4.992 0 0 1 9.878 0l.099.67l.658.156A2.994 2.994 0 0 1 25 28z"/>`),
		g.Group(children),
	)
}