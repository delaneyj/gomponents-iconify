package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.586 17L22 21.586l-5-5a2.002 2.002 0 0 0-2.829 0L9.585 21.17a2.003 2.003 0 0 0 0 2.829l4 4H6v2h10a1 1 0 0 0 .707-1.707L11 22.585L15.585 18l5.708 5.707a1 1 0 0 0 1.414 0L28 18.414Z"/><path fill="currentColor" d="M21.5 17a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5Zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.502 1.502 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="m4 10.598l12-6.462l12.526 6.745l.948-1.762l-13-7a1.004 1.004 0 0 0-.948 0l-13 7A1 1 0 0 0 2 10v20h2Z"/>`),
		g.Group(children),
	)
}