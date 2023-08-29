package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentMultipleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 31H6c-1.1 0-2-.9-2-2V12h2v17h12v2z"/><path fill="currentColor" d="M22 27H10c-1.1 0-2-.9-2-2V8h2v17h12v2z"/><path fill="currentColor" d="M16 16h8v2h-8z"/><path fill="currentColor" d="m27.7 9.3l-7-7c-.2-.2-.4-.3-.7-.3h-6c-1.1 0-2 .9-2 2v17c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V10c0-.3-.1-.5-.3-.7zM20 4.4l5.6 5.6H20V4.4zM26 21H14V4h4v6c0 1.1.9 2 2 2h6v9z"/>`),
		g.Group(children),
	)
}