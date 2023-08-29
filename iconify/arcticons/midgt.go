package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Midgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.315 5.5l-2.452 7.597s-.28 1.155-1.768 1.155H8.823L7.41 19.019h23.533s5.073.19 7.22-5.078s3.218-8.255 3.218-8.255L26.315 5.5Zm-4.63 37l2.452-7.597s.28-1.155 1.768-1.155h13.272l1.412-4.767H17.056s-5.073-.19-7.22 5.078c-2.149 5.268-3.218 8.255-3.218 8.255l15.067.186ZM6.907 24H41.33"/>`),
		g.Group(children),
	)
}