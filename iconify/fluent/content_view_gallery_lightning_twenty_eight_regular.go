package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryLightningTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.75A3.75 3.75 0 0 1 6.75 3h14.5A3.75 3.75 0 0 1 25 6.75V13h-1.5v-2.5h-6v3.28l-.04.078L16 16.94V4.5H6.75A2.25 2.25 0 0 0 4.5 6.75v14.5a2.25 2.25 0 0 0 2.25 2.25h11.196l-.5 1.5H6.75A3.75 3.75 0 0 1 3 21.25V6.75ZM23.5 9V6.75a2.25 2.25 0 0 0-2.25-2.25H17.5V9h6ZM7.75 7.5c-.69 0-1.25.56-1.25 1.25v3.5c0 .69.56 1.25 1.25 1.25h5c.69 0 1.25-.56 1.25-1.25v-3.5c0-.69-.56-1.25-1.25-1.25h-5ZM8 12V9h4.5v3H8Zm-1.5 4.25a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75ZM7.25 19a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5h-6Zm11.567-5h6.329a.5.5 0 0 1 .436.745L23.75 18h1.515a.75.75 0 0 1 .568 1.24l-6.415 7.452c-.523.606-1.5.052-1.246-.707L19.5 22h-4a.5.5 0 0 1-.452-.714l3.317-7a.5.5 0 0 1 .451-.286Z"/>`),
		g.Group(children),
	)
}