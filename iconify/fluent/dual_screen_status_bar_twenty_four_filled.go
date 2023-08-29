package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenStatusBarTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.245 4.004c.967 0 1.75.784 1.75 1.75V18.25a1.75 1.75 0 0 1-1.75 1.75h-7.247c-.087 0-.172-.006-.256-.018V4.023c.083-.012.169-.019.256-.019h7.247Zm-9.247-.002a1.8 1.8 0 0 1 .245.017V19.98a1.8 1.8 0 0 1-.245.017H3.75A1.75 1.75 0 0 1 2 18.247V5.752c0-.967.784-1.75 1.75-1.75h7.248ZM19.75 6.5h-4.502a.75.75 0 0 0-.101 1.493l.101.007h4.502a.75.75 0 0 0 .102-1.493L19.75 6.5Z"/>`),
		g.Group(children),
	)
}