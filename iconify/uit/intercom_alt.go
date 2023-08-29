package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntercomAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14.5a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-1 0v8a.5.5 0 0 0 .5.5zm4 0a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-1 0v8a.5.5 0 0 0 .5.5zM20 1H4a3.003 3.003 0 0 0-3 3v16a3.003 3.003 0 0 0 3 3h16a3.003 3.003 0 0 0 3-3V4a3.003 3.003 0 0 0-3-3zm2 19a2.003 2.003 0 0 1-2 2H4a2.003 2.003 0 0 1-2-2V4a2.003 2.003 0 0 1 2-2h16a2.003 2.003 0 0 1 2 2v16zM6 12.5a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5zm12-5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 1 0V8a.5.5 0 0 0-.5-.5zm-.32 8.291A9.459 9.459 0 0 1 12 17.5a9.447 9.447 0 0 1-5.68-1.71a.5.5 0 0 0-.641.767A10.255 10.255 0 0 0 12 18.5a10.258 10.258 0 0 0 6.321-1.942a.5.5 0 0 0-.641-.767z"/>`),
		g.Group(children),
	)
}