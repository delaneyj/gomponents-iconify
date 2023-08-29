package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M28 19v-4h-2v2H6v-2H4v4h7v4.412l-4 2.225V30h2v-3.188L12.26 25h7.481L23 26.813V30h2v-4.363l-4-2.225V19zm-9 4h-6v-4h6zM0 11h10v2H0zm22 0h10v2H22zm-6-5a5.982 5.982 0 0 0-4.24 1.76l1.413 1.413a3.994 3.994 0 0 1 5.654 0l1.414-1.414A5.981 5.981 0 0 0 16 6z"/><path fill="currentColor" d="m8.932 4.932l1.414 1.414a7.988 7.988 0 0 1 11.308 0l1.414-1.414a9.984 9.984 0 0 0-14.136 0Z"/>`),
		g.Group(children),
	)
}