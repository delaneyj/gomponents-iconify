package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeObjectStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 24a2.98 2.98 0 0 0-2.038.811L16.96 22.41a2.048 2.048 0 0 0 0-.818l4.003-2.403a3.246 3.246 0 1 0-.92-1.779l-4.004 2.402a3 3 0 1 0 0 4.377l4.003 2.403A2.973 2.973 0 0 0 20 27a3 3 0 1 0 3-3Zm0-8a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm-9 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm9 5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M8 28H4a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h7.586A1.986 1.986 0 0 1 13 4.586L16.414 8H28a2.002 2.002 0 0 1 2 2v8h-2v-8H15.586l-4-4H4v20h4Z"/>`),
		g.Group(children),
	)
}