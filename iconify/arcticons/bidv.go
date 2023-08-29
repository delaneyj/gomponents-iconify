package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bidv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.87 40.01a8.884 8.884 0 0 0 9.202-15.022m-9.341-11.251c.337 3.65-.264 5.534-.847 7.271c1.778.452 7.56 1.417 10.188 3.979m-28.31 9.701a8.884 8.884 0 0 0 17.13 4.109m7.812-12.361c-3.367 1.448-5.345 1.46-7.177 1.441c.12 1.832.989 7.628-.635 10.92M10.918 14.868a8.884 8.884 0 0 0 1.384 17.561m14.171 3.611c-2.418-2.754-3.04-4.632-3.589-6.38c-1.704.68-6.949 3.297-10.581 2.77M29.507 7.942a8.884 8.884 0 0 0-16.275 6.744m.946 14.592c1.872-3.15 3.466-4.322 4.959-5.385c-1.173-1.411-5.284-5.59-5.905-9.207m28.606 8.792a8.884 8.884 0 0 0-11.444-13.393"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.81 15.494c3.574.807 5.181 1.96 6.653 3.052c.98-1.552 3.683-6.753 6.932-8.461"/>`),
		g.Group(children),
	)
}