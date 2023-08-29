package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barrier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.703 2a.67.67 0 0 0-.667.673v3.494c0 .211.1.39.25.514a.649.649 0 0 0 .417.159s.305.019.305.028V15h1.007v-2h11.999v2h.994V6.868c0-.009.263-.028.263-.028a.67.67 0 0 0 .667-.673V2.673A.672.672 0 0 0 16.271 2H1.703zm9.308 1H13l-2.01 3H9l2.01-3zM6.729 3h1.713L6.713 6H5l1.729-3zm-1.752 7h2.039l-.963 2H4.016l.96-2zm4 0h2.039l-.963 2H8.016l.96-2zm4 0h2.039l-.963 2h-2.037l.96-2zM2.095 5.71A.624.624 0 0 1 2 5.392V3.61c0-.337.26-.61.582-.61H4L2.582 6a.572.572 0 0 1-.487-.29zM15 7.023V9H3V7h11.042c-.001.01.958.017.958.024zm.357-1.038h-1.31l1.762-3c.166.117.279.29.279.49v1.872c0 .354-.326.638-.731.638z"/>`),
		g.Group(children),
	)
}