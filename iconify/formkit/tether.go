package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tether(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 2h10v2.38H3zm6.5 6.76v5.25h-3V8.76z"/><path fill="currentColor" d="M8 5.75c-3.66 0-6.62.73-6.62 1.62S4.35 8.99 8 8.99s6.62-.73 6.62-1.62S11.65 5.75 8 5.75Zm0 2.5c-3.38 0-6.12-.5-6.12-1.12S4.62 6.01 8 6.01s6.12.5 6.12 1.12S11.38 8.25 8 8.25Z"/><path fill="currentColor" d="M8 7.88c.52 0 1.02-.01 1.5-.04V3.75h-3v4.09c.48.02.98.04 1.5.04Z"/>`),
		g.Group(children),
	)
}