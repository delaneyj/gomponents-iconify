package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21 11.818C21 10.174 22.435 9 24 9s3 1.174 3 2.818v16.364C27 29.826 25.565 31 24 31s-3-1.174-3-2.818V11.818ZM24 11c-.644 0-1 .454-1 .818v16.364c0 .364.356.818 1 .818s1-.454 1-.818V11.818c0-.364-.356-.818-1-.818Zm0 24a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}