package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DottedSixPointedStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="m12.825 10.5l2.309-4a1 1 0 0 1 1.732 0l2.31 4h4.618a1 1 0 0 1 .866 1.5l-2.31 4l2.31 4a1 1 0 0 1-.866 1.5h-4.619l-2.309 4a1 1 0 0 1-1.732 0l-2.31-4H8.207A1 1 0 0 1 7.34 20l2.31-4l-2.31-4a1 1 0 0 1 .866-1.5h4.619Zm1.732 0h2.886L16 8l-1.443 2.5Zm-4.042 4L11.96 12H9.072l1.443 2.5Zm.866 1.5l2.31 4h4.618l2.31-4l-2.31-4h-4.618l-2.31 4Zm-.866 1.5L9.072 20h2.887l-1.444-2.5Zm4.042 4L16 24l1.443-2.5h-2.886Zm8.371-1.5l-1.443-2.5L20.04 20h2.887Zm-1.443-5.5l1.443-2.5h-2.887l1.444 2.5Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}