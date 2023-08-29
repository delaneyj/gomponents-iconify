package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoGlassdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.828 28h10.344a3.43 3.43 0 0 0 3.43-3.43V10.33h-3.43v14.24H7.398a3.43 3.43 0 0 0 3.43 3.43Z"/><path fill="currentColor" d="M21.172 4H10.828a3.43 3.43 0 0 0-3.43 3.43v14.24h3.43V7.43h13.774A3.43 3.43 0 0 0 21.172 4Z"/>`),
		g.Group(children),
	)
}