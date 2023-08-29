package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slideshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 7.08a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5zm5.845-1.137C15.265 8.498 13.616 8.051 12 8c-1.118-.057-1.5.298-1.5 1.08l.001 6c0 5 8.421 3.43 5.165-4.949c1.671-.959 3.076-2.434 3.876-3.412c.411-.608-.028-1.245-.697-.776zM7 2.08a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5zM8 8c-1.616.051-3.265.498-6.845-2.057c-.669-.469-1.108.168-.697.775c.8.979 2.205 2.453 3.876 3.412c-3.256 8.379 5.165 9.949 5.165 4.949l.001-6C9.5 8.298 9.118 7.943 8 8z"/>`),
		g.Group(children),
	)
}