package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1437 753q0 208-87 370.5t-248 254t-369 91.5q-149 0-285-58t-234-156t-156-234T0 736t58-285t156-234T448 61T733 3q286 0 491 192l-199 191Q908 273 733 273q-123 0-227.5 62T340 503.5T279 736t61 232.5T505.5 1137t227.5 62q83 0 152.5-23t114.5-57.5t78.5-78.5t49-83t21.5-74H733V631h692q12 63 12 122zm867-122v210h-209v209h-210V841h-209V631h209V422h210v209h209z"/>`),
		g.Group(children),
	)
}