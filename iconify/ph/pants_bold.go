package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PantsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m227.85 213.52l-22-176A20 20 0 0 0 186 20H70a20 20 0 0 0-19.85 17.52l-22 176A20 20 0 0 0 48 236h40.69a19.94 19.94 0 0 0 19.38-15L128 144l19.91 76.94a20 20 0 0 0 19.4 15.06H208a20 20 0 0 0 19.85-22.48ZM190 104.23A28.05 28.05 0 0 1 176.29 84h11.18ZM182.47 44l2 16H71.53l2-16ZM68.53 84h11.18A28 28 0 0 1 66 104.22ZM85.6 212H52.53l10.14-81.11A52.1 52.1 0 0 0 103.83 84H116v10.47Zm84.81 0L140 94.47V84h12.17a52.09 52.09 0 0 0 41.16 46.88L203.47 212Z"/>`),
		g.Group(children),
	)
}