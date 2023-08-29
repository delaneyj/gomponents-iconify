package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryLightningTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 3h8.25v9.84l-2.36 5.022A1.5 1.5 0 0 0 13.497 20H14.5v1H6.25A3.25 3.25 0 0 1 3 17.75V6.25A3.25 3.25 0 0 1 6.25 3Zm10.067 8H21V8.5h-5v2.534a1.5 1.5 0 0 1 .317-.034ZM16 7V3h1.75A3.25 3.25 0 0 1 21 6.25V7h-5ZM6 13.75c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0-.75.75ZM7.25 6.5C6.56 6.5 6 7.06 6 7.75v2.5c0 .69.56 1.25 1.25 1.25h3.5c.69 0 1.25-.56 1.25-1.25v-2.5c0-.69-.56-1.25-1.25-1.25h-3.5ZM7.5 10V8h3v2h-3ZM6 16.75c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0-.75.75ZM16.317 12h4.828a.5.5 0 0 1 .436.745L19.75 16h1.496a.75.75 0 0 1 .564 1.244l-4.824 5.508c-.504.576-1.442.085-1.257-.657L16.5 19h-3.003a.5.5 0 0 1-.452-.713l2.82-6a.5.5 0 0 1 .452-.287Z"/>`),
		g.Group(children),
	)
}