package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleKeep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.661 10.661C26.661 4.802 21.86 0 16 0S5.339 4.803 5.339 10.661c0 3.401 1.599 6.521 4.26 8.536v10.641h2.161v2.161h8.48v-2.161h2.161V19.197c2.681-2 4.26-5.156 4.26-8.536zM11.74 27.697v-2.099h8.52v2.099zm0-4.26v-2.099h8.52v2.099zm9-5.677l-.48.319v1.119h-8.52v-1.119l-.459-.319a8.48 8.48 0 0 1-3.803-7.099c0-4.697 3.803-8.521 8.521-8.521s8.521 3.803 8.521 8.521c0 2.86-1.401 5.521-3.803 7.099z"/>`),
		g.Group(children),
	)
}