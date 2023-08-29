package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiotechnica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0A11.992 11.992 0 0 0 .015 11.985A12.019 12.019 0 0 0 12 24a12.019 12.019 0 0 0 11.985-12.015A11.992 11.992 0 0 0 12.004 0zm0 .903a11.078 11.078 0 0 1 11.085 11.078c0 6.123-4.958 11.112-11.085 11.112A11.104 11.104 0 0 1 .922 11.985A11.078 11.078 0 0 1 11.996.907zm.087 1.16l-.43 1.252l-5.674 16.063l-.204.604h12.654l-.23-.604L12.524 3.31zm0 2.797l2.007 5.643l-3.024 8.553H7.056zm2.502 7.038l2.532 7.155h-5.09z"/>`),
		g.Group(children),
	)
}