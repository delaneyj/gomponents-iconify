package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryLightningTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 3H6.75A3.75 3.75 0 0 0 3 6.75v14.5A3.75 3.75 0 0 0 6.75 25H17v-2h-1.5a1.5 1.5 0 0 1-1.355-2.142L17 14.83V3Zm8 10h-6.184a1.45 1.45 0 0 0-.316.034V9.5H25V13Zm-6.5-5V3h2.75A3.75 3.75 0 0 1 25 6.75V8h-6.5Zm-12 .75c0-.69.56-1.25 1.25-1.25h5c.69 0 1.25.56 1.25 1.25v3.5c0 .69-.56 1.25-1.25 1.25h-5c-.69 0-1.25-.56-1.25-1.25v-3.5ZM8 9v3h4.5V9H8Zm-.75 8a.75.75 0 0 1 0-1.5h6a.75.75 0 0 1 0 1.5h-6Zm-.75 2.75a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75ZM18.817 14h6.329a.5.5 0 0 1 .436.745L23.75 18h1.515a.75.75 0 0 1 .568 1.24l-6.415 7.452c-.523.606-1.5.052-1.246-.707L19.5 22h-4a.5.5 0 0 1-.452-.714l3.317-7a.5.5 0 0 1 .451-.286Z"/>`),
		g.Group(children),
	)
}