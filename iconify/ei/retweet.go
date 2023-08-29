package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M38 35h-2V17c0-.6-.4-1-1-1H18v-2h17c1.7 0 3 1.3 3 3v18z"/><path fill="currentColor" d="m37 36.5l-6.8-7.8l1.6-1.4l5.2 6.2l5.2-6.2l1.6 1.4zm-5-.5H15c-1.7 0-3-1.3-3-3V15h2v18c0 .6.4 1 1 1h17v2z"/><path fill="currentColor" d="M18.2 22.7L13 16.5l-5.2 6.2l-1.6-1.4l6.8-7.8l6.8 7.8z"/>`),
		g.Group(children),
	)
}