package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viruswar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="5.23" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 10.96V7.95"/><circle cx="24" cy="24.05" r="13.09" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.1" cy="9.63" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.41 14.02l1.94-2.3"/><circle cx="42.54" cy="20.78" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.89 21.77l2.96-.52"/><circle cx="40.3" cy="33.46" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.34 30.59l2.6 1.5"/><circle cx="30.44" cy="41.73" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.48 36.35l1.02 2.82"/><circle cx="17.56" cy="41.73" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.52 36.35l-1.02 2.82"/><circle cx="7.7" cy="33.46" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.66 30.59l-2.6 1.5"/><circle cx="5.46" cy="20.78" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.11 21.77l-2.96-.52"/><circle cx="11.9" cy="9.63" r="2.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.59 14.02l-1.94-2.3"/>`),
		g.Group(children),
	)
}