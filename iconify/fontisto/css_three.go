package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.366 21.423L21.279 0H0l.025.278L1.91 21.426l.016.176L10.559 24h.136l8.653-2.4l.016-.176zM10.658 4.679h6.183l-.186 2.16l-6.1 2.712l.102.488h5.724l-.659 7.551L10.636 19l-5.084-1.408l-.32-3.616h2.093l.168 1.833l.016.178l3.085.82h.133l3.051-.848l.015-.177l.288-3.386l.023-.276H5.057l-.19-2.173l6.309-2.701l-.1-.49h-6.48l-.185-2.08h6.24z"/>`),
		g.Group(children),
	)
}