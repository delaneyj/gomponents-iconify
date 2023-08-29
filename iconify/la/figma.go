package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.5 1A5.506 5.506 0 0 0 7 6.5c0 1.658.74 3.143 1.904 4.152A5.505 5.505 0 0 0 6 15.5c0 1.86.932 3.504 2.35 4.5A5.493 5.493 0 0 0 6 24.5c0 3.033 2.467 5.5 5.5 5.5s5.5-2.467 5.5-5.5V12h3.5c3.033 0 5.5-2.467 5.5-5.5S23.533 1 20.5 1h-8zm0 2h8C22.43 3 24 4.57 24 6.5S22.43 10 20.5 10h-8C10.57 10 9 8.43 9 6.5S10.57 3 12.5 3zm-1 9H15v7h-3.5C9.57 19 8 17.43 8 15.5S9.57 12 11.5 12zm9.5 0a4 4 0 0 0 0 8a4 4 0 0 0 0-8zm-9.5 9H15v3.5c0 1.93-1.57 3.5-3.5 3.5S8 26.43 8 24.5S9.57 21 11.5 21z"/>`),
		g.Group(children),
	)
}