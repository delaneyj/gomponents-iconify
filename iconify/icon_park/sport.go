package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" d="M36 15C38.7614 15 41 12.7614 41 10C41 7.23858 38.7614 5 36 5C33.2386 5 31 7.23858 31 10C31 12.7614 33.2386 15 36 15Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 16.7691L20.0031 13.998L31 19.2466L20.0031 27.4442L31 34.6834L24.0083 43.998"/><path stroke-linecap="round" stroke-linejoin="round" d="M35.3203 21.6434L38.002 23.1018L44.0003 17.4658"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.849 31.5454L13.8793 35.4572L4.00391 40.9964"/></g>`),
		g.Group(children),
	)
}