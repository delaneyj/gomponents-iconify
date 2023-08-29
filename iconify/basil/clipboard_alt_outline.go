package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.65 4.263l.813.101a3.75 3.75 0 0 1 3.287 3.721v10.49a3.641 3.641 0 0 1-3.191 3.613c-3.028.377-6.09.377-9.118 0a3.641 3.641 0 0 1-3.191-3.613V8.085a3.75 3.75 0 0 1 3.287-3.72l.813-.102A2.751 2.751 0 0 1 11 2.25h2a2.75 2.75 0 0 1 2.65 2.013Zm-7.4 1.524l-.528.066A2.25 2.25 0 0 0 5.75 8.085v10.49c0 1.08.805 1.991 1.876 2.125a35.39 35.39 0 0 0 8.747 0a2.141 2.141 0 0 0 1.877-2.125V8.085a2.25 2.25 0 0 0-1.972-2.232l-.528-.066V7a.75.75 0 0 1-.75.75H9A.75.75 0 0 1 8.25 7V5.787ZM9.75 5c0-.69.56-1.25 1.25-1.25h2c.69 0 1.25.56 1.25 1.25v1.25h-4.5V5Z" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 11.75A.75.75 0 0 0 15 11H9a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75Zm-1 3A.75.75 0 0 0 14 14H9a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 .75-.75Z"/>`),
		g.Group(children),
	)
}