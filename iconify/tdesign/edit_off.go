package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.748 2.947a2 2 0 0 1 2.829 0l2.474 2.475a2 2 0 0 1 0 2.829l-2.35 2.35l-2.202 2.202l-1.415-1.414l1.496-1.496l-2.475-2.474l-1.495 1.495L11.195 7.5l2.203-2.203l2.35-2.35Zm-.228 3.057l2.474 2.475l1.643-1.643l-2.475-2.474l-1.642 1.642ZM4 2.586L21.414 20L20 21.414l-6.056-6.056l-4.786 4.786l-6.38 1.076l1.077-6.38l4.785-4.785L2.586 4L4 2.586Zm6.055 8.883L5.72 15.803l-.503 2.977l2.977-.502l4.335-4.334l-2.475-2.475Z"/>`),
		g.Group(children),
	)
}