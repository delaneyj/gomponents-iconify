package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 26H4a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h24a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM4 10v14h24V10Z"/><path fill="currentColor" d="M10 20h11v2H10zm-4-8h2v2H6zm4 0h2v2h-2zm4 0h2v2h-2zm4 0h2v2h-2zM6 20h2v2H6zm0-4h2v2H6zm4 0h2v2h-2zm4 0h2v2h-2zm8-4h4v2h-4zm0 4h4v2h-4zm-4 0h2v2h-2zm5 4h3v2h-3z"/>`),
		g.Group(children),
	)
}