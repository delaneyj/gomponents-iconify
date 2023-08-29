package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M365.7 102.4v146.3H512V102.4H365.7zm109.7 109.7h-73.1V139h73.1v73.1zM0 431.6h146.3V285.3H0v146.3zm36.6-109.7h73.1V395H36.6v-73.1zM0 248.7h146.3V102.4H0v146.3zM36.6 139h73.1v73.1H36.6V139zm146.3 109.7h146.3V102.4H182.9v146.3zM219.4 139h73.1v73.1h-73.1V139zm-36.5 292.6h146.3V285.3H182.9v146.3zm36.5-109.7h73.1V395h-73.1v-73.1zm146.3 109.7H512V285.3H365.7v146.3zm36.6-109.7h73.1V395h-73.1v-73.1z"/>`),
		g.Group(children),
	)
}