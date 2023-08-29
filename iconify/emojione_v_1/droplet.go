package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#57c6e9" d="M35.898 3.677C34.224-.55 29.041-2.087 24.72 5.419s-30.354 42.994.976 56.875c11.538 5.292 25.322-2.938 27.353-12.245c3.627-16.611-11.161-31.25-17.15-46.37"/><ellipse cx="21.11" cy="51.53" fill="#e7e6e6" rx="4.296" ry="6.719" transform="rotate(159.298 21.114 51.532)"/>`),
		g.Group(children),
	)
}