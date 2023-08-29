package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowHookUpRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 11a5 5 0 0 0 5 5h4.5a.5.5 0 0 0 0-1H9a4 4 0 0 1 0-8h5.293l-2.7 2.7a.5.5 0 1 0 .708.706l3.539-3.539a.5.5 0 0 0 .125-.497a.498.498 0 0 0-.135-.247L12.297 2.59a.5.5 0 1 0-.707.707L14.293 6H9a5 5 0 0 0-5 5Z"/>`),
		g.Group(children),
	)
}