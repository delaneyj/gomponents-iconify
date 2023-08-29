package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 0h175c58 0 109 27 141 71c23 30 35 66 35 105s-12 75-35 104c-3 5-7 10-10 15c34 12 63 31 88 58c38 40 61 96 61 155c0 60-23 114-61 155c-41 44-100 72-166 72H0V0zm72 280h108c55-2 99-48 99-104S235 74 180 71H72v209zm0 73v310h158c85-1 153-70 153-155s-68-155-153-155H72z"/>`),
		g.Group(children),
	)
}