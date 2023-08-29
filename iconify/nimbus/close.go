package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Close(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m14.41 3.27l-.82-.94L8 7.17L2.41 2.33l-.82.94L7.05 8l-5.46 4.73l.82.94L8 8.83l5.59 4.84l.82-.94L8.95 8l5.46-4.73z"/>`),
		g.Group(children),
	)
}