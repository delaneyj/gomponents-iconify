package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SupportVectorMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="26" cy="18" r="4" fill="currentColor"/><circle cx="18" cy="26" r="4" fill="currentColor"/><path fill="currentColor" d="M2 28.586L28.586 2L30 3.414L3.414 30zM14 10c-2.206 0-4-1.794-4-4s1.794-4 4-4s4 1.794 4 4s-1.794 4-4 4zm0-6c-1.103 0-2 .897-2 2s.897 2 2 2s2-.897 2-2s-.897-2-2-2zM6 18c-2.206 0-4-1.794-4-4s1.794-4 4-4s4 1.794 4 4s-1.794 4-4 4zm0-6c-1.103 0-2 .897-2 2s.897 2 2 2s2-.897 2-2s-.897-2-2-2z"/>`),
		g.Group(children),
	)
}