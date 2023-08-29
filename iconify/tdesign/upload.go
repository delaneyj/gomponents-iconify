package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2.586L17.914 8.5L16.5 9.914l-3.5-3.5V16h-2V6.414l-3.5 3.5L6.086 8.5L12 2.586ZM4.5 14v5h15v-5h2v7h-19v-7h2Z"/>`),
		g.Group(children),
	)
}