package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncompleteWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 26a12 12 0 0 1 0-24zm3.826-21.236a10.029 10.029 0 0 1 3.242 2.168L22.48 5.52a12.036 12.036 0 0 0-3.89-2.602zM26 14a11.93 11.93 0 0 0-.917-4.59l-1.847.764A9.943 9.943 0 0 1 24 14zm1.38 14h-6.762L24 21.236zM24 18a1 1 0 0 0-.895.553l-5 10A1 1 0 0 0 19 30h10a1 1 0 0 0 .921-1.39l-5.026-10.057A1 1 0 0 0 24 18z"/>`),
		g.Group(children),
	)
}