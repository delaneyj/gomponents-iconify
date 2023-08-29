package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rads(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#9d4bef"/><path fill="#fff" fill-rule="nonzero" d="M11.47 7.661a3.808 3.808 0 1 1 0 7.616a3.808 3.808 0 0 1 0-7.616zm3.807 12.87a3.808 3.808 0 1 1-3.808-3.808a5.253 5.253 0 0 0 5.253-5.253a3.808 3.808 0 1 1 3.808 3.808a5.253 5.253 0 0 0-5.252 5.253zm5.253 3.808a3.808 3.808 0 1 1 0-7.616a3.808 3.808 0 0 1 0 7.616zm0-2.66a1.148 1.148 0 1 0 0-2.296a1.148 1.148 0 0 0 0 2.296zm-9.06 0a1.148 1.148 0 1 0 0-2.296a1.148 1.148 0 0 0 0 2.296zm9.06-9.062a1.148 1.148 0 1 0 0-2.296a1.148 1.148 0 0 0 0 2.296zm-9.06 0a1.148 1.148 0 1 0 0-2.296a1.148 1.148 0 0 0 0 2.296z"/></g>`),
		g.Group(children),
	)
}