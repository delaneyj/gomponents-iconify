package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInstagram0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M34 6H14a8 8 0 0 0-8 8v20a8 8 0 0 0 8 8h20a8 8 0 0 0 8-8V14a8 8 0 0 0-8-8Z"/><path fill="#000" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M24 32a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/><path fill="#000" d="M35 15a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInstagram0)"/>`),
		g.Group(children),
	)
}