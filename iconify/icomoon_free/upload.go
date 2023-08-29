package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 9h2V5h3L8 1L4 5h3zm3-2.25v1.542L14.579 10L8 12.453L1.421 10L6 8.292V6.75L0 9v4l8 3l8-3V9z"/>`),
		g.Group(children),
	)
}