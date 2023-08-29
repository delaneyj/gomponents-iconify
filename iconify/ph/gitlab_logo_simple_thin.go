package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitlabLogoSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M226.27 118.11L206.38 42a7.94 7.94 0 0 0-15.16-.75L172.57 92H83.43L64.78 41.24a7.94 7.94 0 0 0-15.16.76l-19.89 76.11a53.16 53.16 0 0 0 20.44 56.68l73.27 51.76a7.9 7.9 0 0 0 9.12 0l73.27-51.76a53.16 53.16 0 0 0 20.44-56.68Zm-25.05 50.14L128 220l-73.22-51.75a45.11 45.11 0 0 1-17.31-48.11l19.84-76l19.58 53.24a4 4 0 0 0 3.75 2.62h94.72a4 4 0 0 0 3.75-2.62l19.57-53.22l19.85 76a45.11 45.11 0 0 1-17.31 48.09Z"/>`),
		g.Group(children),
	)
}