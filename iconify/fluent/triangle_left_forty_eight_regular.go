package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.435 22.241c-1.239.652-1.249 2.422-.017 3.088l29.501 15.954a1.75 1.75 0 0 0 2.583-1.54V8.256a1.75 1.75 0 0 0-2.566-1.549L8.435 22.241Zm-1.207 5.287c-2.99-1.617-2.966-5.915.042-7.499L36.772 4.495c2.83-1.49 6.23.562 6.23 3.76v31.488c0 3.22-3.44 5.27-6.272 3.739L7.228 27.528Z"/>`),
		g.Group(children),
	)
}