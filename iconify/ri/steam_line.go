package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteamLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.001 4a5 5 0 1 1-.892 9.92l-2.651 1.989a4 4 0 1 1-7.923.068L1.709 14.43l.75-1.854l3.826 1.545a3.994 3.994 0 0 1 3.697-1.592l2.04-3.061A5 5 0 0 1 17.002 4Zm-7.5 10.5c-.464 0-.892.158-1.231.424l1.606.649a1 1 0 0 1-.75 1.854L7.52 16.78a2 2 0 1 0 1.981-2.28Zm3.364-2.69l-.983 1.476c.284.21.54.458.758.735l1.36-1.02a5.027 5.027 0 0 1-1.135-1.191ZM17 6a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm0 1a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}