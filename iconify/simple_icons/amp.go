package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0c6.628 0 12 5.373 12 12s-5.372 12-12 12C5.373 24 0 18.627 0 12S5.373 0 12 0zm-.92 19.278l5.034-8.377a.444.444 0 0 0 .097-.268a.455.455 0 0 0-.455-.455l-2.851.004l.924-5.468l-.927-.003l-5.018 8.367s-.1.183-.1.291c0 .251.204.455.455.455l2.831-.004l-.901 5.458z"/>`),
		g.Group(children),
	)
}