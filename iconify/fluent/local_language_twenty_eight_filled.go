package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLanguageTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.001 2a1 1 0 0 0 0 2H18v2.5c0 .828-.316 1.417-.8 1.817c-.506.417-1.265.683-2.2.683a1 1 0 1 0 0 2c1.295 0 2.537-.367 3.474-1.141C19.43 9.067 20 7.906 20 6.5V3a1 1 0 0 0-1-1h-5.999Zm-2.073 5.615a1 1 0 0 0-1.85 0L2.076 24.619a1 1 0 1 0 1.85.762l1.807-4.39h8.54l1.807 4.39a1 1 0 1 0 1.85-.762l-1.96-4.76a.995.995 0 0 0-.194-.471L10.928 7.615Zm2.521 11.377H6.556l3.447-8.37l3.446 8.37ZM22 2a1 1 0 0 1 1 1v5h2a1 1 0 1 1 0 2h-2v9a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}