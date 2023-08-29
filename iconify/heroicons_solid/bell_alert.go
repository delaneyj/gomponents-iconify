package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M4.214 3.227a.75.75 0 0 0-1.156-.956a8.97 8.97 0 0 0-1.856 3.826a.75.75 0 0 0 1.466.316a7.47 7.47 0 0 1 1.546-3.186Zm12.728-.956a.75.75 0 0 0-1.157.956a7.47 7.47 0 0 1 1.547 3.186a.75.75 0 0 0 1.466-.316a8.971 8.971 0 0 0-1.856-3.826Z"/><path fill-rule="evenodd" d="M10 2a6 6 0 0 0-6 6c0 1.887-.454 3.665-1.257 5.234a.75.75 0 0 0 .515 1.076a32.94 32.94 0 0 0 3.256.508a3.5 3.5 0 0 0 6.972 0a32.933 32.933 0 0 0 3.256-.508a.75.75 0 0 0 .515-1.076A11.448 11.448 0 0 1 16 8a6 6 0 0 0-6-6Zm0 14.5a2 2 0 0 1-1.95-1.557a33.54 33.54 0 0 0 3.9 0A2 2 0 0 1 10 16.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}