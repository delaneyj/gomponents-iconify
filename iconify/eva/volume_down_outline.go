package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaVolumeDownOutline0"><g id="evaVolumeDownOutline1"><path id="evaVolumeDownOutline2" fill="currentColor" d="M20.78 8.37a1 1 0 1 0-1.56 1.26a4 4 0 0 1 0 4.74A1 1 0 0 0 20 16a1 1 0 0 0 .78-.37a6 6 0 0 0 0-7.26Zm-4.31-5.25a1 1 0 0 0-1 0L9 7.57H4a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4A1.06 1.06 0 0 0 16 21a1 1 0 0 0 1-1V4a1 1 0 0 0-.53-.88ZM15 18.1l-5.1-3.5a1 1 0 0 0-.57-.17H5V9.57h4.33a1 1 0 0 0 .57-.17L15 5.9Z"/></g></g>`),
		g.Group(children),
	)
}