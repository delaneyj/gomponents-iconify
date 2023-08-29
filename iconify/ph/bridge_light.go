package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 162h-34V95.28a69.81 69.81 0 0 0 31.75 25.63a6 6 0 1 0 4.5-11.12A57.8 57.8 0 0 1 198 56a6 6 0 0 0-12 0a58 58 0 0 1-116 0a6 6 0 0 0-12 0a57.8 57.8 0 0 1-36.25 53.79a6 6 0 1 0 4.5 11.12A69.81 69.81 0 0 0 58 95.28V162H24a6 6 0 0 0 0 12h34v26a6 6 0 0 0 12 0v-26h116v26a6 6 0 0 0 12 0v-26h34a6 6 0 0 0 0-12Zm-86-38.35V162h-36v-38.35a70.11 70.11 0 0 0 36 0Zm-76-28.5a70.49 70.49 0 0 0 28 24.09V162H70ZM158 162v-42.76a70.49 70.49 0 0 0 28-24.09V162Z"/>`),
		g.Group(children),
	)
}