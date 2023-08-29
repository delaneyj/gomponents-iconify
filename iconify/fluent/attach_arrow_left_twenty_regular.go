package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachArrowLeftTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M4.828 10.486l5.657-5.657a3 3 0 0 1 4.304 4.179c.39.02.77.081 1.135.179a4.001 4.001 0 0 0-6.146-5.065L4.12 9.779a.5.5 0 1 0 .707.707z" fill="currentColor"/><path d="M8.01 15.789l.997-.997c.022.418.09.825.201 1.213l-.49.491a2.5 2.5 0 0 1-3.536-3.535l6.01-6.01a.5.5 0 1 1 .707.706l-6.01 6.01A1.5 1.5 0 1 0 8.01 15.79z" fill="currentColor"/><path d="M19 14.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0zm-2.147.354l.002-.002a.498.498 0 0 0 .145-.349v-.006a.496.496 0 0 0-.147-.35l-2-2a.5.5 0 1 0-.707.707L15.293 14H12.5a.5.5 0 1 0 0 1h2.793l-1.147 1.147a.5.5 0 1 0 .707.707l2-2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}