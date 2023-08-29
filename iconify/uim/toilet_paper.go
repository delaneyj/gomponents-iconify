package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToiletPaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<ellipse cx="7" cy="7.993" fill="currentColor" rx="1" ry="1.25"/><path fill="currentColor" d="M7 2C4.243 2 2 4.691 2 8s2.243 6 5 6s5-2.691 5-6s-2.243-6-5-6Zm0 7.243a1.146 1.146 0 0 1-1-1.25a1.146 1.146 0 0 1 1-1.25a1.146 1.146 0 0 1 1 1.25a1.146 1.146 0 0 1-1 1.25Z" opacity=".5"/><path fill="currentColor" d="M22.76 20.35A7.504 7.504 0 0 1 21 15.459V8c0-3.309-2.243-6-5-6H7c2.757 0 5 2.691 5 6v7.459a9.507 9.507 0 0 0 2.24 6.191A1.001 1.001 0 0 0 15 22h7a1 1 0 0 0 .76-1.65Z" opacity=".25"/><path fill="currentColor" d="M12 8c0 3.309-2.243 6-5 6h5Z"/>`),
		g.Group(children),
	)
}