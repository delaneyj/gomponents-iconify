package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedUmbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.506 6.43a1.475 1.475 0 0 0-2.079.007a1.466 1.466 0 0 0 .01 2.076a1 1 0 1 1-1.414 1.414a3.466 3.466 0 0 1-.01-4.904a3.475 3.475 0 0 1 4.9-.014l.003.003l5.377 5.355A3.136 3.136 0 0 1 16.17 8.49v.02l.02-.02h-.02c.02-1.53 2.24-3.41 3.14-2.51c3.569 3.56 9.49 15.956 9.684 20.022l.462.46a1 1 0 0 1-1.412 1.417l-.447-.446c-4.09-.108-16.553-6.056-20.147-9.633c-.84-.83.96-3.08 2.43-3.07a3.142 3.142 0 0 1 2.028-2.92l-5.4-5.38h-.001Z"/>`),
		g.Group(children),
	)
}