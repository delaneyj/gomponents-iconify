package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lewdfi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="36.424" height="28.952" x="5.788" y="14.048" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.825"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.635 14.048L26.232 5"/><circle cx="32.257" cy="28.524" r="4.825" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.133 28.524h11.905m-11.905-3.619h11.905m-11.905 7.238h11.905"/>`),
		g.Group(children),
	)
}