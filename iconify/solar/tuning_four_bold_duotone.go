package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningFourBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm0 6a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-2 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path d="M13.855 12.75a1.994 1.994 0 0 0 0-1.5H19a.75.75 0 0 1 0 1.5h-5.145Zm-3.71-1.5a1.994 1.994 0 0 0 0 1.5H5a.75.75 0 0 1 0-1.5h5.145Zm1.71 8a1.994 1.994 0 0 1 0 1.5H19a.75.75 0 0 0 0-1.5h-7.145Zm-3.71 0H5a.75.75 0 0 0 0 1.5h3.145a1.994 1.994 0 0 1 0-1.5Zm7.71-14.5a1.993 1.993 0 0 0 0-1.5H19a.75.75 0 0 1 0 1.5h-3.145Zm-3.71 0a1.993 1.993 0 0 1 0-1.5H5a.75.75 0 0 0 0 1.5h7.145Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}