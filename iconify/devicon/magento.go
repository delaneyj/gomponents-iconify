package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#F26322" d="M119.82 31.97v64.01l-15.85 9.12V41.17l-39.62-22.9l-39.64 22.9l.1 63.96l-15.82-9.15V32.02L64.45 0l55.37 31.97zM72.3 105.1l-7.9 4.6l-7.95-4.55V41.17l-15.82 9.15l.03 63.96L64.38 128l23.77-13.72V50.29L72.3 41.14v63.96z"/>`),
		g.Group(children),
	)
}