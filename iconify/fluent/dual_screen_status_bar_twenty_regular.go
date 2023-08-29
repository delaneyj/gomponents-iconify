package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenStatusBarTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6Zm2-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h5.5V5H4Zm12 10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-5.5v10H16Zm-4-9a.5.5 0 0 0 0 1h3.5a.5.5 0 0 0 0-1H12Z"/>`),
		g.Group(children),
	)
}