package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavIconGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h5.219v5.219H0zm9.39 0h5.219v5.219H9.39zm8.608 0h5.219v5.219h-5.219zM0 9.39h5.219v5.219H0zm9.39 0h5.219v5.219H9.39zm8.608 0h5.219v5.219h-5.219zM0 18.781h5.219V24H0zm9.39 0h5.219V24H9.39zm8.608 0h5.219V24h-5.219z"/>`),
		g.Group(children),
	)
}