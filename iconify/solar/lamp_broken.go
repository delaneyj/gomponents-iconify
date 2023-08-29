package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 22h6m-3 0v-7m0-13c2.423 0 3.635 0 4.582.566a4 4 0 0 1 .757.585c.786.775 1.09 1.948 1.7 4.294l.084.324c.828 3.189 1.241 4.783.49 5.903a2.985 2.985 0 0 1-.247.319C18.47 15 16.824 15 13.529 15h-3.058c-3.295 0-4.942 0-5.837-1.01c-.09-.1-.172-.206-.247-.318c-.751-1.12-.337-2.714.49-5.903l.085-.324c.608-2.346.913-3.519 1.699-4.294a4 4 0 0 1 .757-.585c.182-.109.374-.197.582-.268M17.5 15v2"/>`),
		g.Group(children),
	)
}