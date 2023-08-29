package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothConnectLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.311 12l4.343 4.343L12.998 22h-2v-6.686l-4.364 4.364l-1.415-1.415l5.779-5.778v-.97L5.219 5.737l1.415-1.415l4.364 4.364V2h2l5.656 5.657L14.311 12Zm-1.313 1.515v5.657l2.828-2.829l-2.828-2.828Zm0-3.03l2.828-2.828l-2.828-2.829v5.657ZM19.5 13.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-13 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}