package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Setter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="19.781" cy="9.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.124" ry="5.087" transform="rotate(-35.6 19.781 9.6)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.765 43.467c1.908-4.446-2.415-12.517-1.836-16.026c.797-3.023 2.082-9.154-.036-13.821c1.112 1.557 1.264.996.938-.616c.619.703 1.495 2.843 1.586.544c.615 2.567 1.31 1.307 1.738.662a3.679 3.679 0 0 0 3.172 1.575c-6.391 2.586-3.596 8.763-3.668 14.184c.058-4.592 5.087-6.757 6.409-5.258c.232.264-.949 1.673-.808 2.078c.275.775 1.977 2.31 1.315 3.397c-1.72.746-1.08 4.262.854 7.38c1.956-2.806 6.561-6.695 6.58-9.078c-1.775-12.03-6.312-10.892-7.724-14.328c1.716.93 2.06.181 1.662-1.477c.438.101 1.264 2.317 1.191-.399c.55 2.734 1.58-.347 1.521.018c-.072.45.66 1.061 1.202.185c.192-.3.08 2.426.08 2.426c4.172 4.215 6.247 10.946 6.181 12.416c.084.772-1.65 7.054-4.689 12.022a15.23 15.23 0 0 0 1.441 4.15Z"/>`),
		g.Group(children),
	)
}