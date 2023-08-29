package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m15.909 4.561l-4.47-4.47c-.146-.146-.338-.113-.427.073l-.599 1.248l4.175 4.175l1.248-.599c.186-.089.219-.282.073-.427zM9.615 2.115L5.5 2.458c-.273.034-.501.092-.58.449v.001C3.804 8.268 0 13.499 0 13.499l.896.896l4.25-4.25a1.5 1.5 0 1 1 .707.707l-4.25 4.25l.896.896s5.231-3.804 10.591-4.92h.001c.357-.078.415-.306.449-.58l.343-4.115l-4.269-4.269z"/>`),
		g.Group(children),
	)
}