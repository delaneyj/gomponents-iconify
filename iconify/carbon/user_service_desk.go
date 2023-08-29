package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserServiceDesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25.334 11.95l1.206-1.206a1.178 1.178 0 0 1 1.259-.258l1.47.586A1.174 1.174 0 0 1 30 12.15v2.692A1.168 1.168 0 0 1 28.823 16l-.05-.002c-10.296-.64-12.373-9.362-12.766-12.7a1.168 1.168 0 0 1 1.032-1.29A1.15 1.15 0 0 1 17.175 2h2.6a1.163 1.163 0 0 1 1.076.732l.587 1.47a1.163 1.163 0 0 1-.253 1.26l-1.206 1.206s.694 4.7 5.355 5.282zM16 30h-2v-5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v5H2v-5a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM9 10a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z"/>`),
		g.Group(children),
	)
}