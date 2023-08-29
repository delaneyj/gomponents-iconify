package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaHardDriveOutline0"><g id="evaHardDriveOutline1"><g id="evaHardDriveOutline2" fill="currentColor"><path d="m20.79 11.34l-3.34-6.68A3 3 0 0 0 14.76 3H9.24a3 3 0 0 0-2.69 1.66l-3.34 6.68a2 2 0 0 0-.21.9V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.76a2 2 0 0 0-.21-.9ZM8.34 5.55a1 1 0 0 1 .9-.55h5.52a1 1 0 0 1 .9.55L18.38 11H5.62ZM18 19H6a1 1 0 0 1-1-1v-5h14v5a1 1 0 0 1-1 1Z"/><path d="M16 15h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/><circle cx="8" cy="16" r="1"/></g></g></g>`),
		g.Group(children),
	)
}