package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionSignalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 13a8.909 8.909 0 0 1-2.525 6.234l1.413 1.413A10.98 10.98 0 0 0 23.601 4.2l-1.202 1.6A8.932 8.932 0 0 1 26 13Z"/><path fill="currentColor" d="M21 13a5.002 5.002 0 0 1-.902 2.856l1.427 1.428a6.983 6.983 0 0 0-.858-9.501l-1.334 1.49A5.008 5.008 0 0 1 21 13zm9 15.586L3.414 2L2 3.414l3.71 3.71A10.982 10.982 0 0 0 8.4 21.801l1.2-1.6A8.968 8.968 0 0 1 7.172 8.585l2.197 2.197a6.966 6.966 0 0 0 1.964 7.435l1.334-1.49A5.007 5.007 0 0 1 11 13a4.885 4.885 0 0 1 .04-.546l3.96 3.96V30h2V18.414L28.586 30z"/>`),
		g.Group(children),
	)
}