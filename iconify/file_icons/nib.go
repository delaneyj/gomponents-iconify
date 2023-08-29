package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m374.88 0l135.356 135.355l-37.008 37.009L337.873 37.008L374.88 0zM0 492.757l224.546-224.546c-12.117-22.344-2.207-52.409 24.062-61.347c31.372-10.674 62.738 16.725 56.324 49.2c-5.67 28.7-37.12 42.495-61.583 30.269L17.682 512c178.756-95.769 314.017-121.682 376.403-114.301c5.26-43.78 10.448-180.833 52.31-222.571L335.108 63.842c-41.739 41.86-178.79 47.048-222.572 52.31C119.827 177.763 87.783 337.021 0 492.756z"/>`),
		g.Group(children),
	)
}