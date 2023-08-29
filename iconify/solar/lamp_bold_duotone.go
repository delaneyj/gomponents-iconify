package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.968 7.445c.609-2.346.913-3.519 1.7-4.294a4 4 0 0 1 .756-.585C8.372 2 9.584 2 12.007 2s3.634 0 4.582.566a4 4 0 0 1 .757.585c.786.775 1.09 1.948 1.699 4.294l.084.324c.828 3.189 1.242 4.783.49 5.903a2.998 2.998 0 0 1-.247.319c-.285.322-.648.541-1.116.69c-.596.146-1.246.23-1.497.254c-.849.065-1.904.065-3.223.065h-3.059c-3.294 0-4.941 0-5.836-1.01c-.09-.1-.172-.206-.247-.318c-.752-1.12-.338-2.714.49-5.903l.084-.324Z" clip-rule="evenodd"/><path d="M16.759 14.935a.773.773 0 0 0-.003.065v2a.75.75 0 1 0 1.5 0v-2.318c-.596.145-1.246.23-1.497.253Z" opacity=".5"/><path fill-rule="evenodd" d="M11.256 21.25V15h1.5v6.25h-1.5Z" clip-rule="evenodd" opacity=".5"/><path d="M11.256 21.25h-2.25a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5h-3.75Z"/></g>`),
		g.Group(children),
	)
}