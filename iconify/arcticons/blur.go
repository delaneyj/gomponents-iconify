package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.834 37.463V10.537m0 0H27.11a6.724 6.724 0 0 1 6.717 6.731h0A6.724 6.724 0 0 1 27.11 24H15.834"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.669 24H27.11a6.724 6.724 0 0 1 6.717 6.731h0a6.724 6.724 0 0 1-6.717 6.732H15.834M27.11 24H43.5m0 0v-4.683M38.817 24v-4.683"/><circle cx="8.597" cy="24" r="4.097" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}