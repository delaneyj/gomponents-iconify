package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c-5.1-4.813-14-6.766-14-15.271c0-4.717 4.33-16.367 14-23.729c9.67 7.362 14 19.012 14 23.729c0 8.505-8.9 10.459-14 15.271Zm0-39v39"/>`),
		g.Group(children),
	)
}