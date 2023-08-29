package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#006bff" d="M128 3.83L48.72 22.547L36.977 124.17ZM39.464 24.264L0 33.167l35.658 90.604Z"/>`),
		g.Group(children),
	)
}