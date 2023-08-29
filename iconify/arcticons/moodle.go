package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moodle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.8 18.8l-.106 7.86M29.79 24v14.182M16.319 23.409c-.07 4.857 0 9.88 0 14.773m27.182 0V24c-.028-7.434-13.59-7.528-13.71 0a5.419 5.419 0 0 0-2.6-4.727m1.42-1.891l-2.365-2.246l6.5-5.318A58.863 58.863 0 0 0 4.5 18.8h8.51v4.018c5.755 1.312 11.214.902 15.6-5.436Z"/>`),
		g.Group(children),
	)
}