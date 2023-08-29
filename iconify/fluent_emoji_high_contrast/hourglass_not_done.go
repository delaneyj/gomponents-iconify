package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassNotDone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2v3.5c-.19 2.484-1.823 6.52-7 7.348v2.304c5.177.829 6.81 4.864 7 7.348V28a1 1 0 1 1 0 2H7a1 1 0 1 1 0-2v-3.5c.19-2.484 1.823-6.52 7-7.348v-2.304C8.823 14.02 7.19 9.984 7 7.5V4a1 1 0 0 1-1-1Zm4.498 25h11.003c.9 0 1.499-1 .899-1.7c-1.1-1.5-2.8-2.6-4.6-3c-.5-.2-.8-.6-.8-1.1v-7.4c0-.5.3-.9.8-1c3.2-.7 5.6-3.3 6.1-6.6c.1-.6-.4-1.2-1-1.2H9.1c-.6 0-1.1.6-1 1.2c.5 3.3 2.9 5.9 6.1 6.6c.5.1.8.5.8 1v7.4c0 .5-.3.9-.8 1c-1.9.4-3.5 1.5-4.6 3c-.6.8 0 1.799.898 1.8Z"/>`),
		g.Group(children),
	)
}