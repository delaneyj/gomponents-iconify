package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WechatLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.79 187.33a80 80 0 0 0-62.21-114.74a80 80 0 1 0-145.37 66.74l-7.66 26.82A14 14 0 0 0 30 184a13.65 13.65 0 0 0 3.84-.54l26.82-7.66a78.86 78.86 0 0 0 25.77 7.63a80 80 0 0 0 108.89 40.37l26.82 7.66a14 14 0 0 0 17.3-17.3ZM132 152a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm-52 0a79.31 79.31 0 0 0 1.31 14.3a63.34 63.34 0 0 1-16-6.09a8 8 0 0 0-6-.68l-26.41 7.55l7.55-26.41a8 8 0 0 0-.68-6a64 64 0 0 1 111.91-62.24A80.12 80.12 0 0 0 80 152Zm108 0a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}