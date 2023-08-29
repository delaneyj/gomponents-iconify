package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartSpeakerMinimalisticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.24 2.461a4.833 4.833 0 0 0-3.92 4.197l-.016.137a46.932 46.932 0 0 0 0 10.69A4.714 4.714 0 0 0 7.41 21.62l1.016.126c.495.062.742.093.99.118c1.717.18 3.45.18 5.168 0c.248-.025.495-.056.99-.118l.906-.113a4.838 4.838 0 0 0 4.215-4.26l.017-.154a46.818 46.818 0 0 0-.035-10.672a4.701 4.701 0 0 0-3.809-4.066l-.323-.06a24.736 24.736 0 0 0-9.09 0l-.216.04Zm-.815 3.057a.75.75 0 0 1 1.057-.093C8.059 5.908 9.676 6.75 12 6.75s3.941-.842 4.518-1.325a.75.75 0 1 1 .963 1.15C16.62 7.297 14.657 8.25 12 8.25c-2.656 0-4.62-.953-5.482-1.675a.75.75 0 0 1-.093-1.057Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}