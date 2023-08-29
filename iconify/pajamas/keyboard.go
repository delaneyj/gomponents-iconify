package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5ZM2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H2Zm2 8.75a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1-.75-.75ZM11 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8.5-3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-2-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}