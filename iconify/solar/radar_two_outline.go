package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.394 3.065a9.218 9.218 0 0 0-6.938.878a.75.75 0 1 1-.737-1.306a10.718 10.718 0 0 1 8.064-1.021c5.734 1.537 9.138 7.431 7.6 13.166c-1.536 5.735-7.43 9.138-13.165 7.601C3.483 20.847.08 14.952 1.617 9.217a10.74 10.74 0 0 1 1.392-3.115c.574-.875 1.732-.943 2.465-.35l6.998 5.665a.75.75 0 0 1-.944 1.166l-4.04-3.27a5.25 5.25 0 1 0 2.851-2.296a.75.75 0 1 1-.473-1.423a6.75 6.75 0 1 1-3.553 2.768L4.53 6.918a.228.228 0 0 0-.168-.052a.128.128 0 0 0-.099.059a9.24 9.24 0 0 0-1.198 2.68a9.25 9.25 0 1 0 11.33-6.54Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}