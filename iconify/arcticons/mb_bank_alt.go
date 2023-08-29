package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MbBankAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.912 33.884V43.5m0-31.52v9.616M15.618 33.174l-9.145 2.972m29.977-9.742l-9.145 2.972M11.057 20.56l-5.652-7.78m18.527 25.502l-5.652-7.78m4.274-18.222l5.653-7.78M9.679 30l5.652-7.779m18.119-1.337l9.145 2.971m-29.978-9.739l9.146 2.972"/>`),
		g.Group(children),
	)
}