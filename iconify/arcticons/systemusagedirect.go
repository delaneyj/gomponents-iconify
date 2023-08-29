package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Systemusagedirect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 5.21h-39m37.05 31.2V5.21M4.5 36.41h39M6.45 5.21v31.2M33.2 15.64a2.32 2.32 0 1 1-4.64 0v-4.13a2.32 2.32 0 0 1 4.64 0Zm-13.76 7.48a2.32 2.32 0 0 1-4.64 0V19a2.32 2.32 0 0 1 4.64 0ZM33.2 30.6a2.32 2.32 0 0 1-4.64 0v-4.13a2.32 2.32 0 0 1 4.64 0ZM11.33 13.58h17.23M11.33 28.53h17.23m4.64-14.95h3.47M33.2 28.53h3.47m-17.23-7.47h17.23m-21.87 0h-3.47m.6 21.4l6.03-6.05m6.04 0v6.38m12.07-.33l-6.03-6.05"/>`),
		g.Group(children),
	)
}