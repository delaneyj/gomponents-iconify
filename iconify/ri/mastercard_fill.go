package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MastercardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 6.654a6.786 6.786 0 0 1 2.596 5.344A6.786 6.786 0 0 1 12 17.342a6.786 6.786 0 0 1-2.596-5.344A6.786 6.786 0 0 1 12 6.654Zm-.87-.582A7.783 7.783 0 0 0 8.402 12a7.782 7.782 0 0 0 2.728 5.926a6.798 6.798 0 1 1 .003-11.854Zm1.742 11.854A7.783 7.783 0 0 0 15.601 12a7.783 7.783 0 0 0-2.73-5.928a6.798 6.798 0 1 1 .003 11.854Z"/>`),
		g.Group(children),
	)
}