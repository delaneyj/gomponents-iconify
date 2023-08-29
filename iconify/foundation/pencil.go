package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M88.79 29.297L70.702 11.209a2.692 2.692 0 0 0-3.808 0L17.389 60.713l.109.109l-.171-.046l-6.772 25.272l.004.001a2.681 2.681 0 0 0 .642 2.748a2.684 2.684 0 0 0 3.033.531l.002.009l25.027-6.706l-.016-.059l.038.038L88.79 33.105a2.692 2.692 0 0 0 0-3.808zm-69.998 51.85l4.022-15.009l10.988 10.988l-15.01 4.021z"/>`),
		g.Group(children),
	)
}