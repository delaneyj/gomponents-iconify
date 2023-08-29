package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.089 16.71A9 9 0 0 1 8.97 8.326a8.912 8.912 0 0 1 11.941 8.384a.5.5 0 0 0 1 0a10.033 10.033 0 0 0-6.46-9.291a9.981 9.981 0 0 0-11.06 2.944a10.058 10.058 0 0 0-2.3 6.347a.5.5 0 0 0 1 0Z"/><path fill="currentColor" d="M5.985 16.71A6.078 6.078 0 0 1 12 10.7a6.078 6.078 0 0 1 6.015 6.015a.5.5 0 0 0 1 0a7.013 7.013 0 0 0-12.409-4.487a7.151 7.151 0 0 0-1.621 4.482a.5.5 0 0 0 1 0Z"/><path fill="currentColor" d="M8.88 16.71a3.12 3.12 0 0 1 6.24 0a.5.5 0 0 0 1 0a4.119 4.119 0 0 0-7.255-2.669a4.219 4.219 0 0 0-.985 2.669a.5.5 0 0 0 1 0Z"/>`),
		g.Group(children),
	)
}