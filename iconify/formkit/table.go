package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.5 16h-12C.67 16 0 15.33 0 14.5v-13C0 .67.67 0 1.5 0h12c.83 0 1.5.67 1.5 1.5v13c0 .83-.67 1.5-1.5 1.5ZM1.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h12c.28 0 .5-.22.5-.5v-13c0-.28-.22-.5-.5-.5h-12Z"/><path fill="currentColor" d="M7 .62h1v14.75H7z"/><path fill="currentColor" d="M.38 5H14.5v1H.38zM.5 9.99h14v1H.5z"/>`),
		g.Group(children),
	)
}