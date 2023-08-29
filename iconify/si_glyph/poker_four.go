package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PokerFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.372.058H5.627a1.62 1.62 0 0 0-1.613 1.624v12.689a1.62 1.62 0 0 0 1.613 1.624h6.745a1.62 1.62 0 0 0 1.614-1.624V1.682A1.62 1.62 0 0 0 12.372.058zM9.023 11S6.034 9.544 6.034 6.71c0-.979.678-1.776 1.513-1.776c.719 0 1.318.587 1.475 1.374c.15-.787.744-1.374 1.46-1.374c.828 0 1.499.787 1.499 1.76C11.98 9.58 9.023 11 9.023 11z"/>`),
		g.Group(children),
	)
}