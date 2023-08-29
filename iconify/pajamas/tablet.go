package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5H4a.5.5 0 0 0-.5.5v1h9V2a.5.5 0 0 0-.5-.5Zm.5 3h-9V14a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V4.5ZM4 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H4Zm1 12.25a.75.75 0 0 1 .75-.75h4.75a.75.75 0 0 1 0 1.5H5.75a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}