package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PubmedSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M48 32C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48H48zm69.564 64s49.092 11.125 46.596 94.781c0 0 41.47-117.171 204.567 1.649c0 42.788-.315 172.242-.315 223.57c-176.897-149.88-207.385-22.068-207.385-22.068c0-79.856-81.754-70.34-81.754-70.34v-212.65s18.756 1.402 38.291 11.11V96zm86.147 98.283l-24.002 141.35h36.562l11.815-81.38h.373l32.447 81.38h14.633l33.94-81.38h.373l10.314 81.38h36.783l-21.402-141.35h-36.563l-30.388 75.541l-28.696-75.54h-36.19z"/>`),
		g.Group(children),
	)
}