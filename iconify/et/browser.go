package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M41.5 10a.5.5 0 0 0-.5.5v20a.5.5 0 0 1-.5.5h-39a.5.5 0 0 1-.5-.5v-20a.5.5 0 0 0-1 0v20c0 .827.673 1.5 1.5 1.5h39c.827 0 1.5-.673 1.5-1.5v-20a.5.5 0 0 0-.5-.5zm-1-10h-39C.673 0 0 .673 0 1.5v6a.5.5 0 0 0 .5.5h41a.5.5 0 0 0 .5-.5v-6c0-.827-.673-1.5-1.5-1.5zm.5 7H1V1.5a.5.5 0 0 1 .5-.5h39a.5.5 0 0 1 .5.5V7z"/><path d="M17.5 28a.5.5 0 0 0 .5-.5v-16a.5.5 0 0 0-.5-.5h-13a.5.5 0 0 0-.5.5v16a.5.5 0 0 0 .5.5h13zM5 12h12v15H5V12zm17.5 3h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1zm0 5h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1zm0 5h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1z"/><circle cx="4" cy="4" r="1"/><circle cx="8" cy="4" r="1"/><circle cx="12" cy="4" r="1"/></g>`),
		g.Group(children),
	)
}