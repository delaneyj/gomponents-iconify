package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeSocialDistancing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M24 12h-6m2.571-3c0 .317.315.792.632 1.19c.41.513.898.962 1.458 1.304c.42.256.93.502 1.339.502M20.571 15c0-.317.315-.792.632-1.19a5.63 5.63 0 0 1 1.458-1.304c.42-.256.93-.502 1.339-.502M0 12h6M3.428 9c0 .317-.314.792-.632 1.19a5.63 5.63 0 0 1-1.457 1.304c-.42.256-.93.502-1.34.502M3.43 15c0-.317-.315-.792-.633-1.19a5.631 5.631 0 0 0-1.457-1.304c-.42-.256-.93-.502-1.34-.502M14 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24m1.85-19.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}