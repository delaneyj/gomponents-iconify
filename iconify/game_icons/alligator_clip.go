package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlligatorClip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m493.533 193.153l-342.76-178.4l-16.106 16.107c15.92 23.5 10.473 59.73-15.28 85.483c-25.753 25.753-62.812 32.066-86.31 16.106l-15.28 15.28L195.785 490.9l152.383-89.612l-108.608-8.672l54.51-54.51l-91.266-7.022l47.49-47.49l-82.178-6.196l39.644-39.645l-70.617-5.37l35.514-35.513l-61.532-4.543c42.474-7.71 79.37-44.36 87.136-87.135l4.957 61.12l35.103-35.102l5.782 70.203l39.23-39.23l6.607 81.766l47.078-47.08l7.02 91.264l54.51-54.51l8.258 106.545l86.723-147.017z"/>`),
		g.Group(children),
	)
}