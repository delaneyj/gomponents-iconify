package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToiletPaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.4 21.2l-.6-.8c-1.1-1.3-1.7-3-1.7-4.8V7.5c0-3-2.5-5.5-5.5-5.5H6C3.2 2 1 4.7 1 8s2.2 6 5 6h4v1.6c0 2 .7 3.9 2 5.4l.6.8c.1.1.2.2.4.2h9c.1 0 .2 0 .3-.1c.2-.2.3-.5.1-.7zM10 13H8.7c.5-.4.9-.9 1.3-1.4V13zm-4 0c-2.2 0-4-2.2-4-5c0-2.7 1.8-5 4-5h.1C8.2 3 10 5.3 10 8c0 2.8-1.8 5-4 5zm7.2 8l-.5-.6c-1.1-1.3-1.7-3-1.7-4.8V8c0-2.1-.9-3.9-2.2-4.9h5.7C17 3 19 5 19 7.5v8.1c0 2 .7 3.9 1.9 5.4h-7.7zM6 6.2c-.9.1-1.6.9-1.5 1.8c-.1.9.6 1.7 1.5 1.7S7.6 8.9 7.5 8c.1-.9-.6-1.7-1.5-1.8zm0 2.5c-.3 0-.5-.3-.5-.7s.2-.8.5-.8s.5.4.5.8s-.2.7-.5.7z"/>`),
		g.Group(children),
	)
}