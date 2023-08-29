package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeroXbtc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#FF914D" fill-rule="nonzero"/><g fill="#FFF"><path d="M16 27.258c-6.218 0-11.258-5.04-11.258-11.258C4.742 9.782 9.782 4.742 16 4.742c6.218 0 11.258 5.04 11.258 11.258c0 6.218-5.04 11.258-11.258 11.258zm0-.662c5.852 0 10.596-4.744 10.596-10.596S21.852 5.404 16 5.404S5.404 10.148 5.404 16S10.148 26.596 16 26.596z"/><path fill-rule="nonzero" d="M14.09 24.132c.248-.102 1.218-.85 2.155-1.655a51.685 51.685 0 0 0 3.348-3.113c1.026-1.026 1.407-1.467 1.47-1.695c.222-.801-2.205-9.606-2.699-9.795c-.301-.116-2.993 2.123-5.377 4.467c-1.398 1.374-1.785 1.831-1.815 2.13c-.02.231.06.81.209 1.496c.593 2.722 1.672 6.483 2.218 7.718c.229.524.255.547.49.447z"/></g></g>`),
		g.Group(children),
	)
}