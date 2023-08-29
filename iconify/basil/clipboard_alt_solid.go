package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.25 5A2.75 2.75 0 0 1 11 2.25h2A2.75 2.75 0 0 1 15.75 5v2a.75.75 0 0 1-.75.75H9A.75.75 0 0 1 8.25 7V5ZM11 3.75c-.69 0-1.25.56-1.25 1.25v1.25h4.5V5c0-.69-.56-1.25-1.25-1.25h-2Z"/><path d="M6.487 4.929c.126-.06.267.036.266.176L6.75 7A2.25 2.25 0 0 0 9 9.25h6A2.25 2.25 0 0 0 17.25 7V5.104c0-.14.14-.236.267-.175A3.498 3.498 0 0 1 19.5 8.085v10.49a3.39 3.39 0 0 1-2.972 3.365a36.639 36.639 0 0 1-9.056 0A3.391 3.391 0 0 1 4.5 18.575V8.085a3.5 3.5 0 0 1 1.987-3.156ZM15 12a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 9 12h6Zm-1 3a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 9 15h5Z"/></g>`),
		g.Group(children),
	)
}