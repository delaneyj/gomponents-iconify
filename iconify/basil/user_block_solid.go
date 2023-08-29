package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.25 7.5a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Zm-4 9.5A3.75 3.75 0 0 1 6 13.25h.34c.185 0 .369.03.544.086l.866.283a7.251 7.251 0 0 0 2.967.323a.22.22 0 0 1 .223.3a5.983 5.983 0 0 0-.44 2.258c0 1.252.384 2.415 1.04 3.377c.091.134.003.323-.16.328a40.077 40.077 0 0 1-7.84-.5a1.537 1.537 0 0 1-1.29-1.517V17Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5Zm4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5Zm-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}