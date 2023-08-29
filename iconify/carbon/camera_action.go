package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraAction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 23a6 6 0 1 1 6-6a6 6 0 0 1-6 6Zm0-10a4 4 0 1 0 4 4a4 4 0 0 0-4-4Z"/><path fill="currentColor" d="M29 27H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h26a1 1 0 0 1 1 1v20a1 1 0 0 1-1 1ZM4 25h24V7H4Z"/><path fill="currentColor" d="M19 9h7v2h-7z"/><circle cx="12" cy="17" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}