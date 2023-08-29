package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FillDrip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11.3 3.3L9.9 4.7l1.8 1.8l-6.9 6.9a3.017 3.017 0 0 0 0 4.3l.1.1l6.3 6.3c1.2 1.2 3.1 1.2 4.3 0l7.6-7.6l.7-.7l-9.7-9.7l-.8-.8l-.2-.2l-1.8-1.8zm1.8 4.6l7.9 7.9l-2.2 2.2H7.899L6.2 16.3c-.4-.4-.4-1.1 0-1.5l6.9-6.9zM25 19.3l-.8 1.2s-.5.8-1.1 1.7c-.3.5-.5.9-.7 1.4c-.2.5-.4.8-.4 1.4c0 1.6 1.4 3 3 3s3-1.4 3-3c0-.6-.2-1-.4-1.5s-.5-1-.7-1.4c-.5-.9-1.1-1.7-1.1-1.7l-.8-1.1z"/>`),
		g.Group(children),
	)
}