package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OkButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M4.977 13.875a4.89 4.89 0 0 1 9.78 0v4.227a4.89 4.89 0 0 1-9.78 0v-4.227Zm4.89-2.14a2.14 2.14 0 0 0-2.14 2.14v4.227a2.14 2.14 0 0 0 4.28 0v-4.227a2.14 2.14 0 0 0-2.14-2.14Zm8.508-2.751c.76 0 1.375.616 1.375 1.375v3.082c0 .113.137.168.215.086l3.917-4.088a1.375 1.375 0 1 1 1.986 1.903l-3.446 3.597a.125.125 0 0 0-.015.154l3.734 5.804a1.375 1.375 0 1 1-2.313 1.488l-3.345-5.2a.125.125 0 0 0-.195-.019l-.498.52l-.002.002a.13.13 0 0 0-.038.09v3.84a1.375 1.375 0 1 1-2.75 0v-11.26c0-.759.616-1.375 1.375-1.375Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}