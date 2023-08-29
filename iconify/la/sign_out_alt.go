package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOutAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.383 4 4 9.383 4 16s5.383 12 12 12c4.05 0 7.64-2.012 9.813-5.094l-1.625-1.156A9.985 9.985 0 0 1 16 26c-5.535 0-10-4.465-10-10S10.465 6 16 6a9.99 9.99 0 0 1 8.188 4.25l1.625-1.156A11.987 11.987 0 0 0 16 4zm7.344 7.281l-1.438 1.438L24.188 15H12v2h12.188l-2.282 2.281l1.438 1.438l4-4L28.03 16l-.687-.719z"/>`),
		g.Group(children),
	)
}