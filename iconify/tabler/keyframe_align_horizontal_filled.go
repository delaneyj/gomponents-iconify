package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyframeAlignHorizontalFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 6c-.629 0-1.214.301-1.606.807l-2.908 3.748a2.395 2.395 0 0 0-.011 2.876l2.919 3.762c.39.505.977.807 1.606.807c.629 0 1.214-.301 1.606-.807l2.908-3.748a2.395 2.395 0 0 0 .011-2.876l-2.919-3.762A2.032 2.032 0 0 0 12 6zm-7 5a1 1 0 0 1 .117 1.993L5 13H3a1 1 0 0 1-.117-1.993L3 11h2zm16 0a1 1 0 0 1 .117 1.993L21 13h-2a1 1 0 0 1-.117-1.993L19 11h2z"/></g>`),
		g.Group(children),
	)
}