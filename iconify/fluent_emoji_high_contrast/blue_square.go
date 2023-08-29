package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 1h24a3 3 0 0 1 3 3v24a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3ZM3 4v.291L27.709 29H28a1 1 0 0 0 1-1v-.587L4.588 3H4a1 1 0 0 0-1 1Zm0 2.412v2.88L22.709 29h2.878L3 6.412ZM20.587 29L3 11.413v2.878L17.709 29h2.878Zm-5 0L3 16.413v2.878L12.709 29h2.879Zm-5 0L3 21.413v2.878L7.709 29h2.879Zm-5 0L3 26.413V28a1 1 0 0 0 1 1h1.588ZM29 25.291v-2.878L9.588 3h-2.88L29 25.291ZM11.709 3L29 20.291v-2.878L14.588 3h-2.88Zm5 0L29 15.291v-2.879L19.587 3H16.71Zm5 0L29 10.291V7.412L24.587 3H21.71Zm5 0L29 5.291V4a1 1 0 0 0-1-1h-1.291Z"/>`),
		g.Group(children),
	)
}