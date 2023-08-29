package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.575 18.671a9.451 9.451 0 0 1 8.656-4.746a.877.877 0 0 1 .978.865v18.736s.087.567-.977.567c-4.246 0-6.954-1.758-8.659-4.711h-4.626a.501.501 0 0 1-.5-.555l.002-9.605a.502.502 0 0 1 .5-.55Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.806 31.036a6.183 6.183 0 0 1-5.43-3.508a7.745 7.745 0 0 1 0-7.017a6.183 6.183 0 0 1 5.43-3.508v14.033Zm4.166-8.993a4.188 4.188 0 0 1 .097 3.927m2.132-6.122c2.428 2.678 2.518 5.771.245 8.486m2.128-10.815c4.16 4.181 4.25 8.976.247 13.189"/>`),
		g.Group(children),
	)
}