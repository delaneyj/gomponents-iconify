package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeCapture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.657 26.964v.074A5.962 5.962 0 0 1 17.695 33h0a5.962 5.962 0 0 1-5.963-5.963v-6.074A5.962 5.962 0 0 1 17.695 15h0a5.962 5.962 0 0 1 5.962 5.963v.073M36.26 28.5c0 2.485-2 4.5-4.466 4.5h-1.568c-1.634 0-2.958-1.335-2.958-2.981c0-1.647 1.324-2.982 2.958-2.982h6.042"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.393 22.256c1.084-1.086 1.844-1.181 3.937-1.181c2.376 0 3.938 1.045 3.938 3.87V33"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/>`),
		g.Group(children),
	)
}