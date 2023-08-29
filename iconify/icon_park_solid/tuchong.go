package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuchong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTuchong0"><g fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M5 39h38V28h-6v5H11V15h12V9H5v30Z"/><path stroke-linecap="round" d="M43 16v6c-8 0-14-5-14-13h6c0 4 2 7 8 7Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTuchong0)"/>`),
		g.Group(children),
	)
}