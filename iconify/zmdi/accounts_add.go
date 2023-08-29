package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountsAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 149v43h-64v64H64v-64H0v-43h64V85h43v64h64zm213 22q-10 0-19-3q19-28 19-61q0-34-19-61q9-3 19-3q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zm-106.5 0q-26.5 0-45.5-19t-19-45.5t19-45T277.5 43t45 18.5t18.5 45t-18.5 45.5t-45 19zM419 217q37 6 65 22t28 38v43h-64v-43q0-34-29-60zm-142-4q40 0 84 18t44 46v43H149v-43q0-28 44-46t84-18z"/>`),
		g.Group(children),
	)
}