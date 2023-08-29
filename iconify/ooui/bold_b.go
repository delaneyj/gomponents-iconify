package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.9 1c2.4 0 4.2.3 5.4 1c1.3.7 1.9 1.9 1.9 3.6c0 1-.3 1.9-.7 2.6a3 3 0 0 1-2 1.3a4.8 4.8 0 0 1 1.6.7c.4.3.8.8 1.1 1.3a5 5 0 0 1 .4 2.3a4.6 4.6 0 0 1-1.7 3.8A7.6 7.6 0 0 1 10 19H3.3V1h5.6zm.4 7.1c1.1 0 1.9-.1 2.3-.5c.5-.3.7-.9.7-1.5c0-.7-.3-1.2-.8-1.5c-.5-.3-1.3-.5-2.4-.5h-2v4h2.2zm-2.2 3V16h2.5c1.1 0 2-.3 2.4-.7c.5-.5.7-1 .7-1.8a2 2 0 0 0-.7-1.6c-.5-.4-1.3-.6-2.5-.6H7z"/>`),
		g.Group(children),
	)
}