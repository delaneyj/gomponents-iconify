package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFolderTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2a2 2 0 0 1 2 2v7h-.879L9 6.879A3 3 0 0 0 6.879 6H5c-.35 0-.687.06-1 .17V4a2 2 0 0 1 2-2h6Zm3 9c.35 0 .687.06 1 .17V6a2 2 0 0 0-1-1.732V11ZM5 7a2 2 0 0 0-2 2v6.5A2.5 2.5 0 0 0 5.5 18h9a2.5 2.5 0 0 0 2.5-2.5V14a2 2 0 0 0-2-2h-1.879a1 1 0 0 1-.707-.293L8.293 7.586A2 2 0 0 0 6.879 7H5Z"/>`),
		g.Group(children),
	)
}