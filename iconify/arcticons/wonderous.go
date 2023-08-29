package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wonderous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.226 39.987V19.222c.827-19.63 28.721-19.63 29.547 0v20.765"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.438 19.428v10.435L24 22.218l-7.438 7.748V19.43c.31-9.918 14.566-9.918 14.876 0ZM24 43.5V31.929m-5.785 5.786h11.57m-8.058 2.272l4.546-4.442m-4.546 0l4.546 4.442"/>`),
		g.Group(children),
	)
}