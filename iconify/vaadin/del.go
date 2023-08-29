package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Del(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm3 12H1V3h2a4.111 4.111 0 0 1 3.999 4.517c.013.1.02.236.02.374A4.11 4.11 0 0 1 3.005 12zm10-3H9v.012c0 .607.211 1.164.564 1.603c.252.244.601.397.986.397l.074-.002a2.4 2.4 0 0 0 1.518-.631l.708.712a3.4 3.4 0 0 1-2.225.92l-.09.002a2.393 2.393 0 0 1-1.696-.702a3.522 3.522 0 0 1-.84-2.289v-.041c0-.968.344-1.855.915-2.547a2.144 2.144 0 0 1 1.578-.623l.086-.002a2.33 2.33 0 0 1 1.641.672c.47.532.762 1.23.78 1.996v.524zm2 3h-1V3h1v9z"/><path fill="currentColor" d="M3 4H2v7h1c.31 0 3-.12 3-3.5S3.12 4 3 4zm7.49 2.8a1.432 1.432 0 0 0-1.339 1.192L11.93 8a1.734 1.734 0 0 0-.431-.831a1.355 1.355 0 0 0-.934-.371l-.079.002z"/>`),
		g.Group(children),
	)
}