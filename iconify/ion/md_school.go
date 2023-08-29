package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdSchool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M113.5 281.2v85.3L256 448l142.5-81.5v-85.3L256 362.7l-142.5-81.5zM256 64L32 192l224 128 183.3-104.7v147.4H480V192L256 64z" fill="currentColor"/>`),
		g.Group(children),
	)
}