package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookcatalogue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.3 42.5h34.1m-34-2l7-.2l-.5-34.7l-6.3-.1Zm8.6-27l-.3 26.3l6.5-.2l.8-25.9Zm9-5.6l-1.4 32.2l8.7-.1L33 8.4Zm10 6.3l-1.5 26.3l7.8.1l1.4-26.3Zm1 0a8.51 8.51 0 0 1 5.7-3.7M12.9 5.6L14.8 9l.1 4.6m-1.5 26.7l1.3-.5"/>`),
		g.Group(children),
	)
}