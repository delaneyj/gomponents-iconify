package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakingHead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#5d6d74" d="M38.5 42.6c-4.6-1.5-3.3-2.6 1.1-4.3c2.1-.8.8-2.5.9-3.8c0-.6 2.5-.5 2.3-2.9c-.1-1.7-3.8-4.1-4.8-5.1c-.6-.6 1.2-2.2-.1-3.6c-1.7-1.9-2-5.3-3-7.2c0 0 .8-1.2.2-1.9C29.9 8 10.6 8.6 5.6 17.2c-5.5 9.7-5.6 23 5.9 30.2c5.1 3.2-1.4 14.6-1.4 14.6h20.3c0-1.9-2.3-8.9 1.7-8.6c3.4.3 7.7.1 7.3-3.8c-.1-1.2-.2-2.2.6-3.2c.8-.9 2-2.7-1.5-3.8m4.6-1.8L62 43.3v-5zm15.4 16.3l2-4.3l-17.4-9.4zm2-28.3l-2-4.3l-15.4 13.6z"/>`),
		g.Group(children),
	)
}