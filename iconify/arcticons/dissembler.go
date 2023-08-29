package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dissembler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="38" height="38" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 11.97h9v9h-9zm15.03.06h9v9h-9zm-15.05 24v-9h9v9zm15.05-.01v-9h9v9z"/>`),
		g.Group(children),
	)
}