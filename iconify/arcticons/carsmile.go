package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carsmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.65 5.09c.82.82.82 23.2 0 24s-3.2.73-3.78-.34c-.29-.58-.49-5.23-.49-11.67s.2-11.09.49-11.67c.58-1.07 2.86-1.26 3.78-.34Zm-7.95 6.97c.63.63.78 2.14.78 8.72c0 4.6-.2 8.43-.49 9c-.58 1.07-2.85 1.26-3.77.34s-.83-17.39 0-18.21a2.66 2.66 0 0 1 3.48.19Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.26 33.38a12.87 12.87 0 0 0 4.79 3.68a12.11 12.11 0 0 0 6.88 1.6a12.2 12.2 0 0 0 6.3-1.26a15 15 0 0 0 6.1-5c1.41-2.13 4.12-2 4.7.24c.68 2.57-5.47 8.28-10.8 10a24.64 24.64 0 0 1-13 0C10.51 40.25 4.31 32.6 8 31.2c1.65-.58 2.42-.2 4.22 2.18Z"/>`),
		g.Group(children),
	)
}