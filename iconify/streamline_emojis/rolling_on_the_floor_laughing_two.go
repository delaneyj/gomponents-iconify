package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollingOnTheFloorLaughingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="#ffe500" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ebcb00" d="M24 1.5a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 37a18.25 18.25 0 1 1 18.25-18.25A18.25 18.25 0 0 1 24 38.5Z"/><path fill="#fff48c" d="M8.446 14.425A6 1.5-45 1 0 16.93 5.94a6 1.5-45 1 0-8.485 8.485Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 21.5a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#ffaa54" d="M35 9.83c.59.59.27 1.85-.71 2.83S32 14 31.42 13.37s-.26-1.85.71-2.83s2.24-1.29 2.87-.71Z"/><path fill="#ffb0ca" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M35.64 33.14c-5.5 5.5-13.83 6.26-19.4 2a1.41 1.41 0 0 1 .13-2.25a86.93 86.93 0 0 0 19.05-19.02a1.41 1.41 0 0 1 2.25-.13c4.23 5.57 3.47 13.9-2.03 19.4Z"/><path fill="#ff87af" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.9 28.4a25 25 0 0 1 9.48-6.26a16.21 16.21 0 0 1-15.74 15.74a25 25 0 0 1 6.26-9.48Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.09 7L25 10.44a1 1 0 0 0 1.26 1.26l3.41-1.13"/><path fill="#ffaa54" d="M12.33 32.46c.59.59 1.85.27 2.83-.71s1.29-2.24.71-2.83s-1.87-.26-2.87.71s-1.25 2.24-.67 2.83Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m9.53 23.59l3.41-1.13a1 1 0 0 1 1.26 1.26l-1.13 3.41"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.37 32.92a88.21 88.21 0 0 0 10.27-8.78a88.21 88.21 0 0 0 8.78-10.27a1.41 1.41 0 0 1 2.25-.13c.2.26.38.54.56.82A94.07 94.07 0 0 1 28.6 26.1a94.07 94.07 0 0 1-11.54 9.63c-.28-.18-.56-.36-.82-.56a1.41 1.41 0 0 1 .13-2.25Z"/>`),
		g.Group(children),
	)
}