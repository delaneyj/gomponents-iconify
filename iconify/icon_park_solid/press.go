package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Press(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPress0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M22 43c-4.726-1.767-8.668-7.815-10.64-11.357c-.852-1.53-.403-3.408.964-4.502a3.83 3.83 0 0 1 5.1.283L19 29V17.5a2.5 2.5 0 0 1 5 0v6a2.5 2.5 0 0 1 5 0v2a2.5 2.5 0 0 1 5 0v2a2.5 2.5 0 0 1 5 0v7.868c0 1.07-.265 2.128-.882 3.003C37.095 39.82 35.255 42.034 33 43c-3.5 1.5-6.63 1.634-11 0Z"/><path d="M29 12a8 8 0 1 0-15.748 2m0 0c.088.343.199.677.33 1l-.33-1Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPress0)"/>`),
		g.Group(children),
	)
}