package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crossroads(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.586 6.586A1.986 1.986 0 0 0 25.172 6H17V2h-2v10H6.828a1.986 1.986 0 0 0-1.414.586L2 16l3.414 3.414A1.986 1.986 0 0 0 6.828 20H15v10h2V14h8.172a1.986 1.986 0 0 0 1.414-.586L30 10ZM6.828 18l-2-2l2-2H15v4Zm18.344-6H17V8h8.172l2 2Z"/>`),
		g.Group(children),
	)
}