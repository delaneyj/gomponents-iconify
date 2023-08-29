package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftFamilySafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.143 5.036C9.608 5.036 3.5 11.144 3.5 18.679a13.62 13.62 0 0 0 5.293 10.79v.001a7.612 7.612 0 0 1 5.94-12.373h3.236a6.03 6.03 0 0 0 0-12.061h-.826Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.857 5.036c7.535 0 13.643 6.108 13.643 13.643a13.62 13.62 0 0 1-5.293 10.79v.001a7.612 7.612 0 0 0-5.94-12.373h-3.236a6.03 6.03 0 0 1 0-12.061h.826Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.793 29.47l12.285 12.284a4.132 4.132 0 0 0 5.844 0l12.285-12.285"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.566 31.243c3.79 3.791 9.937 3.791 13.727 0l5.439-5.438a9.918 9.918 0 0 0 2.904-7.013v-1.695"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.49 31.047l-5.242-5.242a9.918 9.918 0 0 1-2.905-7.013v-1.695"/>`),
		g.Group(children),
	)
}