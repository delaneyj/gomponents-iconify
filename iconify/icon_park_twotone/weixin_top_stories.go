package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinTopStories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWeixinTopStories0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 4l7.2 7.53L41.32 14L38.4 24l2.92 10l-10.12 2.47L24 44l-7.2-7.53L6.68 34L9.6 24L6.68 14l10.12-2.47L24 4Z"/><path fill="#555" d="m30.977 11.915l.395 7.829L37.954 24l-6.582 4.256l-.395 7.829L24 32.512l-6.977 3.573l-.395-7.829L10.045 24l6.583-4.256l.395-7.829L24 15.488l6.977-3.573Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWeixinTopStories0)"/>`),
		g.Group(children),
	)
}