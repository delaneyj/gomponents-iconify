package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 32.7679V15.2321C44 13.9073 43.1298 12.7428 41.8448 12.4204C38.3687 11.5484 31.1843 10 24 10C16.8157 10 9.63134 11.5484 6.15517 12.4204C4.87016 12.7428 4 13.9073 4 15.2321V32.7679C4 34.0927 4.87016 35.2572 6.15517 35.5796C9.63133 36.4516 16.8157 38 24 38C31.1843 38 38.3687 36.4516 41.8448 35.5796C43.1298 35.2572 44 34.0927 44 32.7679Z"/><path fill="#43CCF8" stroke="#fff" stroke-linejoin="round" d="M22.5735 29.8986C21.9107 30.3625 21 29.8884 21 29.0793V18.9207C21 18.1116 21.9107 17.6375 22.5735 18.1014L29.8297 23.1808C30.3984 23.5789 30.3984 24.4211 29.8297 24.8192L22.5735 29.8986Z"/></g>`),
		g.Group(children),
	)
}