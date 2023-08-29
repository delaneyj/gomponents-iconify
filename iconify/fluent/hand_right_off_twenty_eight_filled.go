package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandRightOffTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m19.283 20.344l-.552.75a11.248 11.248 0 0 0-1.028 1.687l-.695 1.407a3.25 3.25 0 0 1-2.914 1.81H10.5c-1.216 0-2.368-.686-2.855-1.859C6.982 22.545 6 19.77 6 17.25V7.06L2.22 3.28a.75.75 0 1 1 1.06-1.06l22.5 22.5a.75.75 0 0 1-1.06 1.06l-5.437-5.436ZM8 9.061v3.689a.75.75 0 0 0 1.5 0v-2.19L8 9.06Zm13.52 8.247l-.437.593l-4.986-4.986a.75.75 0 0 0 .403-.665V4.5a1 1 0 1 1 2 0v9.861c.545-.312 1.246-.654 1.968-.836c.795-.22 1.684-.3 2.417-.16c.367.07.755.206 1.066.466c.335.28.549.682.549 1.169a.75.75 0 0 1-.347.633l-2.632 1.675ZM15 3v8.818l-2-2V3a1 1 0 1 1 2 0Zm-3.5 1.5v3.818l-2-2V4.5a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}