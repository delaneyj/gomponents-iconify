package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitjournal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.05 22.912L25.088 4.95a1.54 1.54 0 0 0-2.177 0L4.95 22.91a1.54 1.54 0 0 0 0 2.177L22.91 43.05a1.54 1.54 0 0 0 2.177 0L43.05 25.09a1.54 1.54 0 0 0 0-2.177Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.988 20.94a1.797 1.797 0 1 1 1.8-1.797h0a1.798 1.798 0 0 1-1.8 1.797Zm0 15.486a1.797 1.797 0 1 1 1.8-1.798h0a1.798 1.798 0 0 1-1.8 1.798Zm8.972-5.662a1.797 1.797 0 1 1 1.8-1.798a1.798 1.798 0 0 1-1.8 1.797ZM18.251 14.932l3.231 3.227m2.491 2.489l6.712 6.704m-7.697-6.154v11.288"/>`),
		g.Group(children),
	)
}