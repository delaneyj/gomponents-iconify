package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lastpass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.629 6.857c0-.379.304-.686.686-.686c.378 0 .685.312.685.686v10.286a.684.684 0 0 1-.686.686a.69.69 0 0 1-.686-.686V6.857zM2.057 10.286a2.057 2.057 0 1 1 0 4.114a2.057 2.057 0 0 1 0-4.114zm7.543 0a2.057 2.057 0 1 1 0 4.114a2.057 2.057 0 0 1 0-4.114zm7.543 0a2.057 2.057 0 1 1 0 4.114a2.057 2.057 0 0 1 0-4.114z"/>`),
		g.Group(children),
	)
}