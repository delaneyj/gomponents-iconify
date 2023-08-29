package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPlay0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path d="M20 24v-6.928l6 3.464L32 24l-6 3.464l-6 3.464V24Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPlay0)"/>`),
		g.Group(children),
	)
}