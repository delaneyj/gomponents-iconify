package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Histogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M416 896V128h192v768H416zm-288 0V448h192v448H128zm576 0V320h192v576H704z"/>`),
		g.Group(children),
	)
}