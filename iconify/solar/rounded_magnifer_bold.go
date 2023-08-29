package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedMagniferBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.82 19.7a1.769 1.769 0 0 1 1.9-1.917c.189.016.414.084.643.154l.066.02l.06.018c.211.063.42.126.58.212a1.786 1.786 0 0 1 .638 2.55c-.1.151-.255.307-.41.464l-.045.044l-.044.045c-.156.157-.31.313-.46.414a1.754 1.754 0 0 1-2.528-.643a3.321 3.321 0 0 1-.21-.585l-.018-.06l-.02-.067c-.07-.232-.137-.459-.153-.648Z" clip-rule="evenodd"/><path d="M20.128 11.143c0 5.05-4.058 9.143-9.064 9.143S2 16.192 2 11.143C2 6.093 6.058 2 11.064 2s9.064 4.093 9.064 9.143Z"/></g>`),
		g.Group(children),
	)
}