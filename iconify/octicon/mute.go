package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M8 2.81v10.38c0 .67-.81 1-1.28.53L3 10H1c-.55 0-1-.45-1-1V7c0-.55.45-1 1-1h2l3.72-3.72C7.19 1.81 8 2.14 8 2.81zm7.53 3.22l-1.06-1.06l-1.97 1.97l-1.97-1.97l-1.06 1.06L11.44 8L9.47 9.97l1.06 1.06l1.97-1.97l1.97 1.97l1.06-1.06L13.56 8l1.97-1.97z" fill="currentColor"/>`),
		g.Group(children),
	)
}