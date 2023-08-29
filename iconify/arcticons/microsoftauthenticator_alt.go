package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftauthenticatorAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.275 0 11.507-5.888 11.507-12.837V15.508C32.489 15.508 24 13.992 24 13.992s-8.496 1.516-11.507 1.516v15.155C12.493 37.613 22.725 43.5 24 43.5Zm8.014-30.618q.01-.178.012-.356a8.026 8.026 0 0 0-16.052 0q.002.178.012.356"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 21.534a3.142 3.142 0 0 0-1.406 5.952a3.992 3.992 0 0 0-2.583 3.736h7.978a3.992 3.992 0 0 0-2.583-3.736A3.142 3.142 0 0 0 24 21.534Z"/>`),
		g.Group(children),
	)
}