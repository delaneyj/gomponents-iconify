package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkexternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.01 10.49a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l8.49-8.48c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-8.5 8.48c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M14.5 7c-.28 0-.5-.22-.5-.5V2H9.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5v5c0 .28-.22.5-.5.5Zm-3 8H2.49C1.67 15 1 14.33 1 13.51V4.49C1 3.67 1.67 3 2.49 3H7.5c.28 0 .5.22.5.5s-.22.5-.5.5H2.49a.49.49 0 0 0-.49.49v9.02c0 .27.22.49.49.49h9.01c.27 0 .49-.22.49-.49V8.5c0-.28.22-.5.5-.5s.5.22.5.5v5.01c0 .82-.67 1.49-1.49 1.49Z"/>`),
		g.Group(children),
	)
}