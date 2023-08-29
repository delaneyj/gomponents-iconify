package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSReload0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M22 43c-4.726-1.767-8.668-7.815-10.64-11.357c-.852-1.53-.403-3.408.964-4.502a3.83 3.83 0 0 1 5.1.283L19 29V17.5a2.5 2.5 0 0 1 5 0v6a2.5 2.5 0 0 1 5 0v2a2.5 2.5 0 0 1 5 0v2a2.5 2.5 0 0 1 5 0v7.868c0 1.07-.265 2.128-.882 3.003C37.095 39.82 35.255 42.034 33 43c-3.5 1.5-6.63 1.634-11 0Z"/><path stroke-linejoin="round" d="m36 12l-4 4l-4-4"/><path d="M14 13a9 9 0 1 1 17.849 1.651"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSReload0)"/>`),
		g.Group(children),
	)
}