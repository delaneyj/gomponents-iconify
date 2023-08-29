package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1c323a" d="M62.35 46.4c0 4.644-13.664 8.404-30.521 8.404S1.309 51.041 1.309 46.4c0-4.636 13.663-8.396 30.52-8.396S62.35 41.768 62.35 46.4"/><path fill="#b9b7bc" d="M5.794 45.712c.143 4.944 17.505 6.902 10.06 6.02c-6.776-.807-12.12-3.01-12.12-6.02c0-1.876 5.546-5.846 12.356-5.399c7.633.5-10.425.783-10.292 5.399"/>`),
		g.Group(children),
	)
}