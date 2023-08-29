package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardNoAccessSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6Zm0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11Z"/><path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1ZM4 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}