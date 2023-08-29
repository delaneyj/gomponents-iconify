package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.022 15.688a.758.758 0 0 0-.479.172l-5.229 4.328H17.71a.75.75 0 0 0-.75.75v5.719a.754.754 0 0 0 .75.75h.016l4.952-.088l4.842 4.395a.75.75 0 0 0 .504.193a.762.762 0 0 0 .305-.064a.75.75 0 0 0 .445-.686V16.438a.751.751 0 0 0-.432-.68a.779.779 0 0 0-.32-.07Zm-20.136-.539a17.358 17.358 0 0 1 14.979-8.646m.275 34.993A17.357 17.357 0 0 1 8.16 32.85m29.706-17.701a17.312 17.312 0 0 1 0 17.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.357 16.527a14.586 14.586 0 0 1 12.585-7.265m-.001 29.067a14.58 14.58 0 0 1-12.583-7.262m25.003-14.541a14.546 14.546 0 0 1 .002 14.54M12.816 17.925a11.646 11.646 0 0 1 10.05-5.801m-.001 23.346a11.643 11.643 0 0 1-10.048-5.799m20.226-11.674a11.624 11.624 0 0 1-.002 11.602"/>`),
		g.Group(children),
	)
}