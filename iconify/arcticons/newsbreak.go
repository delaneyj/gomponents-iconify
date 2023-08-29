package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newsbreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.832 23.697l-9.027-12.938l-12.345 4.408l12.962 24.557l14.24-7.671m-20.867-4.884L6.765 40.552c-.48.583-1.42.115-1.243-.62l5.938-24.765"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.422 39.724L41.27 7.537c.342-.654 1.336-.325 1.22.404l-3.828 24.112"/>`),
		g.Group(children),
	)
}