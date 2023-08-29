package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.682 4.524a2.794 2.794 0 1 0 0 5.587a2.794 2.794 0 0 0 0-5.587ZM7.365 7.317a4.317 4.317 0 1 1 8.635 0a4.317 4.317 0 0 1-8.635 0Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M8.127 8.08H3.556V6.555h4.571v1.523Z" clip-rule="evenodd"/><path fill="currentColor" d="M4.063 7.317a1.524 1.524 0 1 1-3.047 0a1.524 1.524 0 0 1 3.047 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M2.54 6.81a.508.508 0 1 0 0 1.015a.508.508 0 0 0 0-1.015ZM0 7.317a2.54 2.54 0 1 1 5.08 0a2.54 2.54 0 0 1-5.08 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}