package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banknotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.938 4.75L3.155 9H15l11.313-1.938L26.655 9h2.032l-.75-4.25zM2 10v16h28V10H2zm4.938 2h18.125a2.426 2.426 0 0 0-.063.5a2.5 2.5 0 0 0 2.5 2.5c.173 0 .337-.03.5-.063v6.126a2.426 2.426 0 0 0-.5-.063a2.5 2.5 0 0 0-2.5 2.5c0 .173.03.337.063.5H6.938c.033-.163.062-.327.062-.5A2.5 2.5 0 0 0 4.5 21c-.173 0-.337.03-.5.063v-6.125c.163.033.327.062.5.062A2.5 2.5 0 0 0 7 12.5c0-.173-.03-.337-.063-.5zM16 13c-2.75 0-5 2.25-5 5s2.25 5 5 5s5-2.25 5-5s-2.25-5-5-5zm0 2c1.67 0 3 1.33 3 3s-1.33 3-3 3s-3-1.33-3-3s1.33-3 3-3z"/>`),
		g.Group(children),
	)
}