package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fuck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 276q0 58 36.5 102t93.5 51q68 8 118-36t50-111v-64q0-13-11-28t-31-15v107q0 47-35 79t-83 28q-41-7-68.5-39.5T45 276V152q-17 0-29.5 12.5T3 195v81zm192-145v85h42v-85q0-22-21-22t-21 22zM131 24v192h42V24q0-21-21-21t-21 21zM67 131v85h42v-85q0-22-21-22t-21 22z"/>`),
		g.Group(children),
	)
}