package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WavePulseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.974 18a1.446 1.446 0 0 1-1.259-.972l-1.843-4.145c-.115-.26-.262-.378-.349-.378H2.562a.5.5 0 1 1 0-1h2.961a1.444 1.444 0 0 1 1.263.972l1.839 4.145c.116.261.258.378.349.378c.088 0 .229-.113.344-.368L13.7 6.956A1.423 1.423 0 0 1 14.958 6a1.449 1.449 0 0 1 1.26.975l1.839 4.151c.11.249.259.379.349.379h3.028a.5.5 0 0 1 0 1H18.41a1.444 1.444 0 0 1-1.263-.975l-1.839-4.151c-.116-.261-.259-.378-.35-.379c-.088 0-.229.114-.344.368l-4.385 9.676A1.437 1.437 0 0 1 8.974 18Z"/>`),
		g.Group(children),
	)
}