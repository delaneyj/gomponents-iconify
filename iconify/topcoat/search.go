package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M1 17.838c0 8.747 7.131 15.827 15.94 15.827c8.796 0 15.938-7.08 15.938-15.827S25.736 2 16.94 2C8.131 2 1 9.091 1 17.838zm5.051 0c0-5.979 4.868-10.817 10.89-10.817c6.01 0 10.888 4.839 10.888 10.817c0 5.979-4.878 10.818-10.888 10.818c-6.022 0-10.89-4.84-10.89-10.818zm22.111 14.523l6.855 7.809c1.104 1.102 1.816 1.111 2.938 0l2.201-2.181c1.082-1.081 1.149-1.778 0-2.921l-7.896-6.775l-4.098 4.068z"/>`),
		g.Group(children),
	)
}