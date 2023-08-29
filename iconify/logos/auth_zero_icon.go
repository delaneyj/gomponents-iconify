package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuthZeroIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M220.412 0h-92.415l28.562 89.006h92.416l-74.77 53.077l28.57 89.511c48.128-35.06 63.854-88.12 46.208-142.588L220.413 0ZM7.018 89.006h92.416L127.997 0H35.589L7.019 89.006c-17.655 54.468-1.92 107.529 46.207 142.588l28.563-89.51l-74.77-53.078Zm46.208 142.588l74.77 52.97l74.77-52.97l-74.77-53.847l-74.77 53.847Z"/>`),
		g.Group(children),
	)
}