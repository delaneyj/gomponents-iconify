package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frontapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosFrontapp0" x1="12.519%" x2="88.228%" y1="85.213%" y2="10.023%"><stop offset="0%" stop-color="#FF0057" stop-opacity=".16"/><stop offset="86.135%" stop-color="#FF0057"/></linearGradient></defs><path fill="#001B38" d="M0 60.854C0 27.245 27.245 0 60.854 0h195.143v86.6c0 16.804-13.623 30.427-30.427 30.427h-79.762c-15.805.25-28.565 13.033-28.781 28.846v189.41c0 16.804-13.622 30.427-30.427 30.427H0V60.854Z"/><circle cx="147.013" cy="147.015" r="78.993" fill="url(#logosFrontapp0)" transform="rotate(90 147.013 147.015)"/><circle cx="147.013" cy="147.015" r="78.993" fill="url(#logosFrontapp0)" opacity=".5" transform="rotate(90 147.013 147.015)"/>`),
		g.Group(children),
	)
}