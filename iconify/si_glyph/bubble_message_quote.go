package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessageQuote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.009.042C3.615.042.052 2.826.052 6.263c0 3.136 2.972 5.722 6.832 6.151l-2.116 3.57l6.183-3.945c2.937-.914 5.018-3.154 5.018-5.776c-.001-3.437-3.566-6.221-7.96-6.221zM7.051 7c0 1.333-1.655 2.062-3.084 2.062V7.945c.721 0 1.741-.354 1.741-.945H3.967V3.989h3.084V7zm5 0c0 1.333-1.655 2.062-3.084 2.062V7.945c.721 0 1.741-.354 1.741-.945H8.967V3.989h3.084V7z"/>`),
		g.Group(children),
	)
}