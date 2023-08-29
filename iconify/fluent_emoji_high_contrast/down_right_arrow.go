package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m10.293 8.707l-1.586 1.586a1 1 0 0 0 0 1.414l7.586 7.586l.018.018c.13.123.176.345.05.472l-1.494 1.493c-.628.628-.186 1.702.702 1.707l6.441.033a1 1 0 0 0 1.006-1.005l-.033-6.442c-.005-.888-1.08-1.33-1.707-.702l-1.493 1.493c-.127.127-.349.081-.472-.049a.49.49 0 0 0-.007-.007l-.011-.011l-7.586-7.586a1 1 0 0 0-1.414 0Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}