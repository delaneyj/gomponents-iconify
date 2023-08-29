package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialTwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M80 32l-32 80v304h96v64h64l64-64h80l112-112V32H80zm176 240h-48V143h48v129zm112 0h-48V143h48v129z" fill="currentColor"/>`),
		g.Group(children),
	)
}