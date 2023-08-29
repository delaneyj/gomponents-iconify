package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteCane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.268 1.732a2.5 2.5 0 0 0-3.536 0L3.405 25.06a3 3 0 1 0 3.536 3.536L26.035 9.5c.174 2.085.988 3.661 1.965 3.661c1.105 0 2-2.014 2-4.5c0-1.022-.151-1.964-.407-2.72l.675-.674a2.5 2.5 0 0 0 0-3.536Zm-7.622 6.914l.707.708l-12.5 12.5l-.707-.708l12.5-12.5Zm-17.75 17.75l.708.707l-.583.583l-.023.378a1 1 0 1 1-1.062-1.062l.378-.023l.582-.583ZM28.631 6.905c.075.506.119 1.11.119 1.757c0 1.795-.336 3.25-.75 3.25s-.75-1.455-.75-3.25c0-.129.002-.256.005-.382l1.376-1.375Z"/>`),
		g.Group(children),
	)
}