package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NameBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 5.5c.137-.098.377-.224.5-.224c.11 0 .33.126.458.2l.042.024c.076.043.207.124.27.194l3.486 3.978a1 1 0 0 0 1.488 0l2.974-3.311c.37-.412 1.007-.448 1.39-.048A13.953 13.953 0 0 1 30 16c0 7.732-6.268 14-14 14S2 23.732 2 16c0-3.759 1.481-7.172 3.892-9.687c.383-.4 1.02-.364 1.39.048l2.974 3.31a1 1 0 0 0 1.488 0l3.49-3.977c.045-.05.125-.102.193-.145c.027-.018.053-.034.073-.049ZM7 16a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H7Z"/>`),
		g.Group(children),
	)
}