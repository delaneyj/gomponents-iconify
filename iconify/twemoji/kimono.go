package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kimono(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M22 0h3l5 2.808l5 23.307L26 27l-6-16.423M14 0h-3L6 2.808L1 26.115L10 27l6-16.423"/><path fill="#3B88C3" d="M11 5h14v31H11z"/><path fill="#269" d="M17.985 8.761L11 14h2l5.094-5.094z"/><path fill="#88C9F9" d="M25 20s-1.167 1-7 1s-7-1-7-1v-6h14v6z"/><path fill="#FFF" d="M25 18s-1.167 1-7 1s-7-1-7-1v-2h14v2zm0-18l-7 9l-7-9z"/><path fill="#99AAB5" d="M20.333 0L18 3l-2.333-3z"/><path fill="#269" d="M25 25V4l2 23zm-14 0V4L9 27z"/>`),
		g.Group(children),
	)
}