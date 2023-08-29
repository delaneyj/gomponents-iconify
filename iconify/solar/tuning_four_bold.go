package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningFourBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm0 6a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-2 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm5.25-6a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75ZM14 19.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5h-5ZM10.75 4a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 .75-.75ZM5 11.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5H5ZM4.25 20a.75.75 0 0 1 .75-.75h1a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75ZM19 3.25a.75.75 0 0 1 0 1.5h-1a.75.75 0 0 1 0-1.5h1Z"/>`),
		g.Group(children),
	)
}