package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecureBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.628 18.36a13.052 13.052 0 0 1 21.668-4.574m2.692 4.82a55.524 55.524 0 0 1 1.927 23.875M15.61 7.542a17.918 17.918 0 0 0-7.734 7.898m28.14-5.146a13.517 13.517 0 0 1 3.37 4.371M5.5 10.269a22.679 22.679 0 0 1 4.037-4.326m28.55-.092a22.686 22.686 0 0 1 3.988 4.194m-31.22 13.632a96.496 96.496 0 0 1 .384 10.073m-.324 5.99a48.473 48.473 0 0 1-.35 2.74m4.508 0a30.706 30.706 0 0 0 .579-4.468m18.12-.016a79.007 79.007 0 0 1-.33 4.484m-13.86-2.418c.05-.692.09-1.387.138-2.07m9.357 2.209a96.41 96.41 0 0 1-.233 2.279m-4.391-4.448c-.025.887-.066 1.77-.127 2.634c-.04.583-.154 1.255-.215 1.815"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.365 17.209a4.407 4.407 0 0 1 8.815 0v2.362"/><rect width="16.899" height="15.778" x="15.323" y="19.739" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.567"/><path fill="none" stroke="currentColor" stroke-miterlimit="7" d="M19.365 17.209v2.53"/><circle cx="23.772" cy="27.01" r="2.255" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}