package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagPanama(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EEE" d="M18 18V5H4a4 4 0 0 0-4 4v9h18zm0 0v13h14a4 4 0 0 0 4-4v-9H18z"/><path fill="#005293" d="M18 18v13H4a4 4 0 0 1-4-4v-9h18z"/><path fill="#D21034" d="M18 18V5h14a4 4 0 0 1 4 4v9H18z"/><path fill="#005293" d="M9.674 10.573L9 8.5l-.674 2.073H6.147l1.763 1.281l-.673 2.073L9 12.646l1.763 1.281l-.673-2.073l1.763-1.281z"/><path fill="#D21034" d="M25.91 24.854l-.673 2.073L27 25.646l1.763 1.281l-.673-2.073l1.763-1.281h-2.179L27 21.5l-.674 2.073h-2.179z"/>`),
		g.Group(children),
	)
}