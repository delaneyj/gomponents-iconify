package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleselection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M894 447q-15-86-67-159l92-92q88 112 105 251H894zM577 130V0q139 17 251 105l-92 92q-73-52-159-67zm-381-25Q308 17 447 0v130q-86 15-159 67zm1 183q-52 73-67 159H0q17-139 105-251zm0 448l-92 92Q17 716 0 577h130q15 86 67 159zm250 158v130q-139-17-251-105l92-92q73 52 159 67zm381 25q-112 88-251 105V894q86-15 159-67zm66-342h130q-17 139-105 251l-92-92q52-73 67-159z"/>`),
		g.Group(children),
	)
}