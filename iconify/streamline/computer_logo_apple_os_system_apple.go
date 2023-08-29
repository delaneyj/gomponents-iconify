package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerLogoAppleOsSystemApple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.18 7.5a2.49 2.49 0 0 1 1.74-2.37a2.91 2.91 0 0 0-4.22-1a1 1 0 0 1-1 0a3.09 3.09 0 0 0-4.41 1.21a5.13 5.13 0 0 0-.54 3.36a7.25 7.25 0 0 0 1.94 4.21a2.09 2.09 0 0 0 2.7.15h0a1.32 1.32 0 0 1 1.57 0a2.06 2.06 0 0 0 2.66-.06a6.58 6.58 0 0 0 1.7-3a2.48 2.48 0 0 1-2.14-2.5Zm-2-5.5L9.68.5"/>`),
		g.Group(children),
	)
}