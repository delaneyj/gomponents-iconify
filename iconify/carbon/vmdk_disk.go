package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmdkDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="10.5" cy="24.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M18 16.414L19.414 15L23 18.585L21.585 20z"/><circle cx="16" cy="13" r="2" fill="currentColor"/><path fill="currentColor" d="M16 6a7 7 0 0 0 0 14v-2a5 5 0 1 1 5-5h2a7 7 0 0 0-7-7Z"/><path fill="currentColor" d="M26 2H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm0 26H6V4h20Z"/>`),
		g.Group(children),
	)
}