package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentArrowRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 5.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L13.707 6H16.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708l-2 2ZM14.5 11c1.33 0 2.55-.472 3.5-1.257v2.533c0 1.418-1.164 2.566-2.6 2.566h-4.59l-4.011 2.961a1.009 1.009 0 0 1-1.4-.199a.978.978 0 0 1-.199-.59v-2.172h-.6c-1.436 0-2.6-1.149-2.6-2.566v-6.71C2 4.149 3.164 3 4.6 3h5a5.5 5.5 0 0 0 4.9 8Z"/>`),
		g.Group(children),
	)
}