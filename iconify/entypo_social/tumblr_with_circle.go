package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrWithCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 .4C4.698.4.4 4.698.4 10s4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6S15.302.4 10 .4zm2.577 13.741a5.508 5.508 0 0 1-1.066.395a4.543 4.543 0 0 1-1.031.113c-.42 0-.791-.055-1.114-.162a2.373 2.373 0 0 1-.826-.459a1.651 1.651 0 0 1-.474-.633c-.088-.225-.132-.549-.132-.973V9.16H6.918V7.846c.359-.119.67-.289.927-.512c.257-.221.464-.486.619-.797c.156-.31.263-.707.322-1.185h1.307v2.35h2.18V9.16h-2.18v2.385c0 .539.028.885.085 1.037a.7.7 0 0 0 .315.367c.204.123.437.185.697.185c.466 0 .928-.154 1.388-.461v1.468z"/>`),
		g.Group(children),
	)
}