package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M362.005 149.115s-7 55.77-79 83.36v-24.69c-2.76-1-4.63 7.88-7.26 9.15h-39.49c-2.63-1.27-4.5-10.11-7.26-9.15v24.69c-72-27.59-79-83.36-79-83.36c-60.71 67.68-121.41 80-121.41 80c102.53-16.11 101.36 44.89 101.36 44.89c69.71-11.91 65.64 36.31 65.64 36.31c60.63-6.9 60.41 52.57 60.41 52.57s-.22-59.47 60.41-52.57c0 0-4.07-48.22 65.64-36.31c0 0-1.16-61 101.37-44.88c.02.01-60.69-12.33-121.41-80.01z"/>`),
		g.Group(children),
	)
}