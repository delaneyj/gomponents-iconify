package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Techmino(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.112 6.325l-1.769-.474L4.5 16.462l1.769.474m9.711-8.771l1.768.474l-2.843 10.611l-1.769-.474m10.852-8.465l-1.769-.474l-2.843 10.611l1.768.474m9.711-8.771l1.769.474l-2.843 10.611l-1.769-.474m-8.024 2.476l-1.769-.474l-2.843 10.611l1.769.474m9.711-8.771l1.768.474l-2.843 10.611l-1.769-.474m10.852-8.465l-1.769-.474l-2.843 10.611l1.768.474m9.711-8.771l1.769.474l-2.843 10.611l-1.769-.474"/>`),
		g.Group(children),
	)
}