package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosPartlysunny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M208 96h16v54h-16z" fill="currentColor"/><path d="M64 240h55v16H64z" fill="currentColor"/><path d="M107.5 149.1l11-11.1 31.4 31.6-11 11.1z" fill="currentColor"/><path d="M287.4 179.7l-11-11.1 31.3-31.6 11.1 11.1z" fill="currentColor"/><path d="M120.5 352.5l-11-11.1 31.4-31.6 11 11.1z" fill="currentColor"/><path d="M299 224c-43.3 0-78.3 35.2-78.3 78.5 0 2.6.1 5.2.4 7.8-26.4 2.3-47.1 25.5-47.1 52.6 0 28.6 23.2 53.1 51.7 53.1h157.7c35.7 0 64.6-29.9 64.6-65.7 0-35.8-28.9-65.3-64.6-65.3-2.7 0-5.4 0-8 .3-7.9-35-39.1-61.3-76.4-61.3z" fill="currentColor"/><path d="M264.7 196.3c-12.9-14.3-31.4-23.3-52-23.3-38.7 0-70.3 31.7-70.3 70.7 0 27.4 15.6 51.2 38.3 62.9v-.1l.1.1c4.5-12.1 11.4-19.8 22.6-25.6.2-.1.4-.2.7-.4.2-.1.5-.2.7-.3-.2-2.3-.3-4.7-.3-7-.3-31.9 30.2-70.5 61-75.4.1-.1.3-.1.4-.2-.4-.4-.8-.9-1.2-1.4z" fill="currentColor"/>`),
		g.Group(children),
	)
}