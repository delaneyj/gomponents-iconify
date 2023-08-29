package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryLightningTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 7H6v2h3V7ZM3 6a3 3 0 0 1 3-3h6v4.783a1.465 1.465 0 0 0-.04.079l-2.416 5.14A.495.495 0 0 0 9.5 13h-4a.5.5 0 0 0 0 1h3.583a1.5 1.5 0 0 0 1.414 2H12v.884l-.029.116H6a3 3 0 0 1-3-3V6Zm14 1h-4V3h1a3 3 0 0 1 3 3v1ZM5 7v2a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1Zm.5 4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Zm7.817-3h4.828a.5.5 0 0 1 .436.745L16.75 12h1.496a.75.75 0 0 1 .564 1.244l-4.824 5.508c-.504.576-1.442.085-1.257-.657L13.5 15h-3.003a.5.5 0 0 1-.452-.713l2.82-6A.5.5 0 0 1 13.317 8Z"/>`),
		g.Group(children),
	)
}