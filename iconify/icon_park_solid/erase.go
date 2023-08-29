package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Erase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSErase0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 42h40"/><path fill="#fff" d="M31 4L7 28l6 6h8l20-20L31 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSErase0)"/>`),
		g.Group(children),
	)
}