package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><circle cx="12" cy="12" r="4"/><path d="M4 24a4 4 0 0 0 4-4v-4H4a4 4 0 1 0 0 8Zm0-8h4V8H4a4 4 0 1 0 0 8Zm0-8h4V0H4a4 4 0 1 0 0 8Zm8 0H8V0h4a4 4 0 1 1 0 8Z"/></g>`),
		g.Group(children),
	)
}