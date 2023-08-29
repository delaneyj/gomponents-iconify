package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M9.754 9A.75.75 0 0 0 9 9.754l.037 7.293a.75.75 0 0 0 1.28.526l1.917-1.916c.127-.127.348-.082.472.049l.018.018l7.586 7.586a1 1 0 0 0 1.414 0l1.586-1.586a1 1 0 0 0 0-1.414l-7.586-7.586a.771.771 0 0 0-.018-.018c-.13-.124-.176-.345-.05-.472l1.917-1.916a.75.75 0 0 0-.526-1.28L9.754 9Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}