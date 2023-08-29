package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TmuxSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1.75 1.5a.25.25 0 0 0-.25.25v12.5c0 .138.112.25.25.25h5.5v-13h-5.5Zm7 0v5.75h5.75v-5.5a.25.25 0 0 0-.25-.25h-5.5Zm5.75 7.25H8.75v5.75h5.5a.25.25 0 0 0 .25-.25v-5.5ZM0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v12.5A1.75 1.75 0 0 1 14.25 16H1.75A1.75 1.75 0 0 1 0 14.25V1.75Z"/>`),
		g.Group(children),
	)
}