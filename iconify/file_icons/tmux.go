package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tmux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M272 22.4h192.005C490.512 22.4 512 43.92 512 70.42V240H272V22.4zm-32 0H47.995C21.51 22.4 0 43.9 0 70.42V489.6h240V22.4zM272 272v217.6h240V272H272z"/>`),
		g.Group(children),
	)
}