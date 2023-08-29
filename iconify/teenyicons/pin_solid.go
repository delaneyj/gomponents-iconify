package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 6.496a1.5 1.5 0 0 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.496A6.499 6.499 0 0 1 7.5 0C11.089 0 14 2.909 14 6.496c0 2.674-1.338 4.793-2.772 6.225a10.865 10.865 0 0 1-2.115 1.654c-.322.19-.623.34-.885.442c-.247.098-.506.174-.728.174c-.222 0-.481-.076-.728-.174a6.453 6.453 0 0 1-.885-.442a10.868 10.868 0 0 1-2.115-1.654C2.338 11.289 1 9.17 1 6.496Zm6.5-2.499a2.5 2.5 0 0 0-2.5 2.5a2.5 2.5 0 0 0 5 0a2.5 2.5 0 0 0-2.5-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}