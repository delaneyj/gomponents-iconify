package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmZOsAiControlInterface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="12" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M23 25H9c-1.103 0-2-.897-2-2v-4c0-1.103.897-2 2-2h14c1.103 0 2 .897 2 2v4c0 1.103-.897 2-2 2ZM9 19v4h14v-4H9Z"/><circle cx="12" cy="11" r="1" fill="currentColor"/><path fill="currentColor" d="M23 15H9c-1.103 0-2-.897-2-2V9c0-1.103.897-2 2-2h14c1.103 0 2 .897 2 2v4c0 1.103-.897 2-2 2ZM9 9v4h14V9H9Z"/><path fill="currentColor" d="M28 30H4c-1.103 0-2-.897-2-2V16h2v12h24V4H16V2h12c1.103 0 2 .897 2 2v24c0 1.103-.897 2-2 2ZM12 0v5h-2V0zM3.414 2L6.95 5.536L5.536 6.95L2 3.414zM0 10h5v2H0z"/>`),
		g.Group(children),
	)
}