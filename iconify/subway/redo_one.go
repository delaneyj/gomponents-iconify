package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M389.6 289.4c0 73.8-59.8 133.6-133.6 133.6c-73.7 0-133.6-59.8-133.6-133.6S182.2 155.8 256 155.8v66.8l113.9-111.3L256 0v66.8c-122.9 0-222.6 99.7-222.6 222.6C33.4 412.3 133.1 512 256 512c122.9 0 222.6-99.7 222.6-222.6h-89z"/>`),
		g.Group(children),
	)
}