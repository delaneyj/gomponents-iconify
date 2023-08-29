package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.413 28.09h-1.009c-1.051 0-1.904-.86-1.904-1.92s.853-1.918 1.904-1.918h3.888"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.224 21.175c.698-.699 1.187-.76 2.534-.76c1.53 0 2.534.673 2.534 2.49v5.185m12.317 0H26.6a1.911 1.911 0 0 1-1.904-1.92c0-1.059.853-1.918 1.904-1.918h3.888"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.42 21.175c.698-.699 1.187-.76 2.534-.76c1.53 0 2.534.673 2.534 2.49v5.185M17 20.415v4.674a3 3 0 1 0 6.003 0v-7.576m9.49 2.902c2.17-.264 2.75 1.731 2.129 2.034c-1.46.71-1.377 1.854-1.377 3.35c0 1.513.962 2.688 2.557 2.688c1.77 0 2.69-1.412 2.69-2.884c0-1.23.008-3.106-1.215-3.258c1.012-.738 1.223-1.635 1.223-1.93m-8.012 0h-2.534m-12.662 0h-2.534"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}