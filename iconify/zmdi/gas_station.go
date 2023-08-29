package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M336 90q16 16 16 38v203q0 22-15.5 37.5T299 384t-38-15.5t-16-37.5V224h-32v160H0V43q0-18 12.5-30.5T43 0h128q17 0 29.5 12.5T213 43v149h22q17 0 29.5 12.5T277 235v96q0 8 6.5 14.5t15 6.5t15-6.5T320 331V177q-11 4-21 4q-22 0-38-15.5T245 128q0-17 9.5-30.5T280 78l-45-45l22-22l80 79h-1zm-165 59V43H43v106h128zm127.5 0q8.5 0 15-6t6.5-15t-6.5-15t-15-6t-15 6t-6.5 15t6.5 15t15 6z"/>`),
		g.Group(children),
	)
}