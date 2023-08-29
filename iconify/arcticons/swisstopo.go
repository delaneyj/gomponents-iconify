package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swisstopo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.836 40.73v-7.155m0 11.925c7.065-2.715 6.66-8.154 6.66-11.4v-2.46a25.214 25.214 0 0 0-6.66-1.14a25.214 25.214 0 0 0-6.66 1.14v2.46c0 3.246-.405 8.685 6.66 11.4Zm-3.578-8.347h7.155"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.54 42.743a21.503 21.503 0 1 1 10.959-18.97m.001.173a21.46 21.46 0 0 1-1.265 7.338"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.82 9.772a9.554 9.554 0 0 0-9.555 9.555c0 7.477 7.314 16.49 9.207 18.698a.577.577 0 0 0 .814.064a.568.568 0 0 0 .068-.068c1.861-2.216 9.02-11.222 9.02-18.694a9.554 9.554 0 0 0-9.554-9.555Zm0 13.125a3.57 3.57 0 1 1 3.57-3.57v0a3.57 3.57 0 0 1-3.57 3.57Z"/>`),
		g.Group(children),
	)
}