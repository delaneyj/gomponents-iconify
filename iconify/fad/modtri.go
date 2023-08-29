package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modtri(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M35.996 208c-2.207 0-3.077-1.532-1.935-3.434l91.878-153.132c1.138-1.896 2.98-1.902 4.122 0l91.878 153.132c1.138 1.896.272 3.434-1.935 3.434h-8.008c-2.207 0-4.922-1.54-6.054-3.421L130.058 78.42c-1.137-1.89-2.984-1.881-4.116 0L50.058 204.58c-1.138 1.889-3.848 3.42-6.054 3.42h-8.008z"/>`),
		g.Group(children),
	)
}