package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleEditTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M12 3.5a8.5 8.5 0 0 0-.705 16.971l-.234.936a2.117 2.117 0 0 0-.064.543C5.945 21.447 2 17.184 2 12C2 6.477 6.477 2 12 2c5.27 0 9.589 4.077 9.972 9.25a3.293 3.293 0 0 0-1.257-.25h-.002c-.09 0-.18.004-.27.011A8.501 8.501 0 0 0 12 3.5Z"/><path d="M20.715 12h-.002c-.585 0-1.17.223-1.615.67l-5.902 5.902a2.684 2.684 0 0 0-.707 1.247l-.458 1.831a1.087 1.087 0 0 0 1.319 1.318l1.83-.457a2.684 2.684 0 0 0 1.248-.707l5.902-5.902A2.285 2.285 0 0 0 20.715 12Z"/></g>`),
		g.Group(children),
	)
}