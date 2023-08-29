package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M6.01245 34.005V8.03613C6.01245 6.93156 6.90788 6.03613 8.01245 6.03613H40.0001C41.1046 6.03613 42.0001 6.93156 42.0001 8.03613V34.005C42.0001 35.6618 40.6569 37.005 39.0001 37.005H37.0049V37.9999C37.0049 40.2091 35.2141 41.9999 33.0049 41.9999H33.0045C30.7954 41.9999 29.0045 40.2091 29.0045 37.9999V37.005H19.008V38.0016C19.008 40.2098 17.2178 41.9999 15.0096 41.9999C12.8014 41.9999 11.0113 40.2098 11.0113 38.0016V37.005H9.01245C7.3556 37.005 6.01245 35.6618 6.01245 34.005Z"/><path stroke-linecap="round" d="M42 23H6"/><path fill="#2F88FF" stroke-linecap="round" d="M34 13H14V23H34V13Z"/><path stroke-linecap="round" d="M14 30H16"/><path stroke-linecap="round" d="M32 30H34"/></g>`),
		g.Group(children),
	)
}