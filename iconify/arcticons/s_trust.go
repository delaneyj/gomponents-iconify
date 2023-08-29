package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STrust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.693 0 15.274-7.782 15.274-16.966V6.504C35.267 6.503 24 4.5 24 4.5S12.723 6.503 8.726 6.503v20.03C8.726 35.719 22.307 43.5 24 43.5Z"/><rect width="19.325" height="16.718" x="14.338" y="16.767" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.746"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.845 22.339h14.817m-19.324 5.573h14.817"/><circle cx="23.987" cy="11.656" r="3.21" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}