package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ereolen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.445 24.512s5.706-.942 7.447-2.742a2.467 2.467 0 0 0 .167-3.365c-1.432-1.856-5.702-1.369-6.494.236c-1.388 2.812-2.718 8.755-1.542 10.247c1 1.268 5.2.968 7.677-.123c3-1.32 6.65-3.778 7.352-6.98a8.706 8.706 0 0 0-3.499-9.05c-3.404-2.36-8.898-2.282-12.588-.4c-3.858 1.968-5.62 5.93-6.273 10.211a12.957 12.957 0 0 0 3.333 10.65a12.168 12.168 0 0 0 6.978 3.636c3.15.515 8.973-.927 8.973-.927"/>`),
		g.Group(children),
	)
}