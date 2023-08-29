package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scaphold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#00ACA1" d="M206.306-.001V18.07H0v73.185h17.669v21.083H0v48.791h224.076v21.082h-17.77v21.986H0v69.271h51.502v-19.276H256v-71.981h-16.263v-21.082H256v-48.791H33.331V91.255H51.2V68.969H256V-.001z"/>`),
		g.Group(children),
	)
}